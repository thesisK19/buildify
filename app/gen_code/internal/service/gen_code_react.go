package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/gen_code/api"
	"github.com/thesisK19/buildify/app/gen_code/internal/constant"
	"github.com/thesisK19/buildify/app/gen_code/internal/dto"
	"github.com/thesisK19/buildify/app/gen_code/internal/util"

	"github.com/iancoleman/strcase"
	"github.com/scylladb/go-set/strset"
)

func (s *Service) GenReactSourceCode(ctx context.Context, request *api.GenReactSourceCodeRequest) (*api.GenReactSourceCodeResponse, error) {
	logger := ctxlogrus.Extract(ctx)
	logger.Info("GenReactSourceCode begin")

	if len(request.Nodes) == 0 || len(request.Pages) == 0 {
		return nil, nil
	}
	return s.doGenReactSourceCode(ctx, request)
}

func (s *Service) doGenReactSourceCode(ctx context.Context, request *api.GenReactSourceCodeRequest) (*api.GenReactSourceCodeResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "doGenReactSourceCode")

	mapPagePathToPageInfo, err := getMapPagePathToPageInfo(ctx, request)
	if err != nil {
		logger.WithError(err).Error("Failed to getMapPagePathToPageInfo")
		return nil, err
	}

	rootDirName := strconv.FormatInt(time.Now().Unix(), constant.BASE_DECIMAL)
	rootDirPath := fmt.Sprintf("%s/%s", constant.EXPORT_DIR, rootDirName)
	outputZipPath := fmt.Sprintf("%s/%s.zip", constant.EXPORT_DIR, rootDirName)

	err = setUpDir(ctx, rootDirPath, request.Pages)
	if err != nil {
		logger.WithError(err).Error("Failed to setUpDir")
		return nil, err
	}

	for _, pageInfo := range mapPagePathToPageInfo {
		err = genPage(ctx, rootDirPath, pageInfo)
		if err != nil {
			logger.WithError(err).Error("Failed to genPage")
			return nil, err
		}
	}

	// index pages
	err = genIndexPages(ctx, rootDirPath, request.Pages)
	if err != nil {
		logger.WithError(err).Error("Failed to genIndexPages")
		return nil, err
	}

	// routing
	err = genRoutes(ctx, rootDirPath, request.Pages)
	if err != nil {
		logger.WithError(err).Error("Failed to genRoutes")
		return nil, err
	}

	// format code
	formatCode(rootDirPath)

	// create zip file
	err = util.ZipDir(ctx, rootDirPath, outputZipPath)
	if err != nil {
		logger.WithError(err).Error("Failed to ZipDir")
		return nil, err
	}

	// upload File
	url, err := util.UploadFile(ctx, outputZipPath, outputZipPath)
	if err != nil {
		logger.WithError(err).Error("Failed to UploadFile")
		return nil, err
	}

	// Remove the input directory after zipping.
	os.RemoveAll(outputZipPath)
	os.RemoveAll(rootDirPath)

	// TODO: delete remote file after 10 mins

	return &api.GenReactSourceCodeResponse{
		Url: *url,
	}, nil
}

func formatCode(rootDirPath string) {
	// npx prettier --write .
	command := exec.Command("prettier", "--write", rootDirPath)
	command.Stderr = os.Stderr
	// Run the command
	command.Run()
}

func genPage(ctx context.Context, rootDirPath string, pageInfo *dto.PageInfo) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "genPage")

	components, mapIDToReactElements, propsByIds := getReactElementInfoFromNodes(pageInfo.Nodes)

	reactElementString := mergeReactElements(pageInfo.RootID, mapIDToReactElements)

	// create props
	propsString := fmt.Sprintf("export const %s = {%s}", constant.PROPS_BY_ID, strings.Join(propsByIds, ","))

	propsFilePath := fmt.Sprintf("%s/%s/%s/%s", rootDirPath, constant.PAGES_DIR, pageInfo.Name, constant.PROPS_TSX)
	err := util.WriteFile(ctx, propsFilePath, []byte(propsString))
	if err != nil {
		logger.WithError(err).Error("Failed to WriteFile")
		return err
	}

	// create page
	content := fmt.Sprintf(`
		import React, { FC, ReactElement } from "react"
   		import { %s } from "src/components"
    	import { %s } from "./props"
	
    	const %s: FC = (): ReactElement => (
		%s
		)
		
		export default %s`,
		strings.Join(components, ","), constant.PROPS_BY_ID, pageInfo.Name, reactElementString, pageInfo.Name)

	indexFilePath := fmt.Sprintf("%s/%s/%s/%s", rootDirPath, constant.PAGES_DIR, pageInfo.Name, constant.INDEX_TSX)
	err = util.WriteFile(ctx, indexFilePath, []byte(content))
	if err != nil {
		logger.WithError(err).Error("Failed to WriteFile")
		return err
	}

	return nil
}

func genIndexPages(ctx context.Context, rootDirPath string, pages []*api.Page) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "genIndexPages")

	var (
		exportPages []string
		importPages []string
	)

	for _, page := range pages {
		importPages = append(importPages, fmt.Sprintf(`import %s from "./%s"`, page.Name, page.Name))
		exportPages = append(exportPages, page.Name)
	}

	content := fmt.Sprintf("%s \n\n export {%s}", strings.Join(importPages, "\n"), strings.Join(exportPages, ","))

	filePath := fmt.Sprintf(`%s/%s/%s`, rootDirPath, constant.PAGES_DIR, constant.INDEX_TS)

	err := util.WriteFile(ctx, filePath, []byte(content))
	if err != nil {
		logger.WithError(err).Error("Failed to WriteFile")
		return err
	}
	return nil
}

func genRoutes(ctx context.Context, rootDirPath string, pages []*api.Page) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "genRoutes")

	var (
		importPages []string
		routes      []string
	)
	for _, page := range pages {
		importPages = append(importPages, page.Name)
		routes = append(routes, fmt.Sprintf(`
		{
			exact: true,
			path: "%s",
			Page: %s,
		}`, page.Path, page.Name))
	}

	content := fmt.Sprintf(`
		import React from "react"
		import {%s} from "../pages"

		export type Route = {
			exact: boolean
			path: string
			Page: React.ElementType
		}

		export const ROUTES: Route[] = [
			%s
		]
	`, strings.Join(importPages, ", "), strings.Join(routes, ","))

	filePath := fmt.Sprintf(`%s/%s/%s`, rootDirPath, constant.ROUTES_DIR, constant.CONFIG_TSX)
	err := util.WriteFile(ctx, filePath, []byte(content))
	if err != nil {
		logger.WithError(err).Error("Failed to WriteFile")
		return err
	}
	return nil
}
func getReactElementInfoFromNodes(nodes []*dto.Node) ([]string, map[string]*dto.ReactElement, []string) {
	components := strset.New()
	mapIDToReactElements := map[string]*dto.ReactElement{}
	propsByIds := []string{}

	for _, node := range nodes {
		reactElement := genReactElementFromNode(node)
		mapIDToReactElements[node.ID] = reactElement
		components.Add(reactElement.Component)
		propsByIds = append(propsByIds, fmt.Sprintf(`"%s": %s`, reactElement.ID, reactElement.Props))
	}

	// sort prop by id
	sort.Slice(propsByIds, func(i, j int) bool {
		return propsByIds[i] < propsByIds[j]
	})

	return components.List(), mapIDToReactElements, propsByIds
}

func genReactElementFromNode(node *dto.Node) *dto.ReactElement {
	elementString := fmt.Sprintf(`<%s {...%s["%s"]}>%s</%s>`, node.ComponentType, constant.PROPS_BY_ID, node.ID, constant.KEY_CHILDREN, node.ComponentType)

	return &dto.ReactElement{
		ID:            node.ID,
		Props:         node.Props,
		Component:     node.ComponentType,
		ElementString: elementString,
		Children:      node.Children,
	}
}
func mergeReactElements(ID string, mapIDToReactElements map[string]*dto.ReactElement) string {
	reactElement, ok := mapIDToReactElements[ID]
	if !ok {
		return ""
	}

	if len(reactElement.Children) == 0 {
		reactElement.ElementString = fmt.Sprintf(`<%s {...%s["%s"]}/>`, reactElement.Component, constant.PROPS_BY_ID, reactElement.ID)
		return reactElement.ElementString
	}

	childrenString := ""
	for _, childID := range reactElement.Children {
		childrenString += mergeReactElements(childID, mapIDToReactElements)
	}

	reactElement.ElementString = strings.Replace(reactElement.ElementString, constant.KEY_CHILDREN, childrenString, 1)

	return reactElement.ElementString
}

func setUpDir(ctx context.Context, rootDirPath string, pages []*api.Page) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "setUpDir")

	err := util.CopyDirRecursively(ctx, constant.REACT_JS_BASE_DIR, rootDirPath)
	if err != nil {
		logger.WithError(err).Error("Failed to CopyDirRecursively")
		return err
	}

	for _, page := range pages {
		err = util.CreateDir(ctx, fmt.Sprintf("%s/%s/%s", rootDirPath, constant.PAGES_DIR, page.Name))
		if err != nil {
			logger.WithError(err).Error("Failed to CreateDir")
			return err
		}
	}

	return nil
}

func getMapPagePathToPageInfo(ctx context.Context, request *api.GenReactSourceCodeRequest) (map[string]*dto.PageInfo, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "getMapPagePathToPageInfo")

	mapPagePathToPageInfo := make(map[string]*dto.PageInfo)
	mapPagePathToPageName := make(map[string]string)

	for _, page := range request.Pages {
		page.Name = strcase.ToCamel(page.Name)
		mapPagePathToPageName[page.Path] = page.Name
	}

	for _, node := range request.Nodes {
		pagePath := node.PagePath
		if pagePath == constant.INVALID_PAGE_PATH {
			err := fmt.Errorf("invalid page path, node=%v", node)
			logger.WithError(err).Error("Failed to CreateDir")
			return nil, err
		}

		_, ok := mapPagePathToPageInfo[pagePath]
		if !ok {
			pageName := mapPagePathToPageName[pagePath]

			mapPagePathToPageInfo[pagePath] = &dto.PageInfo{
				RootID: "",
				Path:   pagePath,
				Name:   pageName,
				Nodes:  []*dto.Node{},
			}

		}

		if strings.HasPrefix(node.Id, constant.ROOT_ID_PREFIX) {
			mapPagePathToPageInfo[pagePath].RootID = node.Id
		}
		mapPagePathToPageInfo[pagePath].Nodes = append(mapPagePathToPageInfo[pagePath].Nodes, &dto.Node{
			ID:            node.GetId(),
			ComponentType: node.GetType(),
			Props:         node.GetProps(),
			Children:      node.GetChildren(),
		})
	}

	return mapPagePathToPageInfo, nil
}