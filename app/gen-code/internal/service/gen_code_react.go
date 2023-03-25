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

	"github.com/thesisK19/buildify/app/gen-code/api"
	"github.com/thesisK19/buildify/app/gen-code/internal/consts"
	"github.com/thesisK19/buildify/app/gen-code/internal/dto"
	"github.com/thesisK19/buildify/app/gen-code/internal/utils"

	"github.com/iancoleman/strcase"
	"github.com/scylladb/go-set/strset"
)

func (s *Service) GenReactSourceCode(ctx context.Context, request *api.GenReactSourceCodeRequest) (*api.GenReactSourceCodeResponse, error) {
	fmt.Println("gen nek")
	if len(request.Nodes) == 0 || len(request.Pages) == 0 {
		return nil, nil
	}
	return s.doGenReactSourceCode(ctx, request)
}

func (s *Service) doGenReactSourceCode(ctx context.Context, request *api.GenReactSourceCodeRequest) (*api.GenReactSourceCodeResponse, error) {
	mapPagePathToPageInfo := getMapPagePathToPageInfo(request)

	rootDirName := strconv.FormatInt(time.Now().Unix(), consts.BASE_DECIMAL)
	err := setUpDir(rootDirName, request.Pages)
	if err != nil {
		return nil, err
	}

	for _, pageInfo := range mapPagePathToPageInfo {
		err = genPage(rootDirName, pageInfo)
		if err != nil {
			return nil, err
		}
	}

	// index pages
	genIndexPages(rootDirName, request.Pages)

	// routing
	genRoutes(rootDirName, request.Pages)

	// format code
	formatCode(rootDirName)

	return &api.GenReactSourceCodeResponse{
		Code:    "OK",
		Message: "OK",
	}, nil
}

func formatCode(rootDirName string) {
	// npx prettier --write .
	command := exec.Command("npx", "prettier", "--write", fmt.Sprintf("%s/%s", consts.EXPORT_DIR, rootDirName))
	command.Stderr = os.Stderr
	// Run the command
	command.Run()
}

func genPage(rootDirName string, pageInfo *dto.PageInfo) error {
	components, mapIDToReactElements, propsByIds := getReactElementInfoFromNodes(pageInfo.Nodes)

	reactElementString := mergeReactElements(pageInfo.RootID, mapIDToReactElements)

	// create props
	propsString := fmt.Sprintf("export const %s = {%s}", consts.PROPS_BY_ID, strings.Join(propsByIds, ","))

	path := fmt.Sprintf("%s/%s/%s/props", rootDirName, consts.PAGES_DIR, pageInfo.Name)
	err := utils.WriteFile(utils.GetFileNameExport(path, "tsx"), []byte(propsString))
	if err != nil {
		return nil
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
		strings.Join(components, ","), consts.PROPS_BY_ID, pageInfo.Name, reactElementString, pageInfo.Name)

	path = fmt.Sprintf("%s/%s/%s/%s", rootDirName, consts.PAGES_DIR, pageInfo.Name, consts.INDEX)
	err = utils.WriteFile(utils.GetFileNameExport(path, "tsx"), []byte(content))
	if err != nil {
		return nil
	}

	return nil
}

func genIndexPages(rootDirName string, pages []*api.Page) error {
	var (
		exportPages []string
		importPages []string
	)

	for _, page := range pages {
		importPages = append(importPages, fmt.Sprintf(`import %s from "./%s"`, page.Name, page.Name))
		exportPages = append(exportPages, page.Name)
	}

	content := fmt.Sprintf("%s \n\n export {%s}", strings.Join(importPages, "\n"), strings.Join(exportPages, ","))

	path := fmt.Sprintf(`%s/%s/%s`, rootDirName, consts.PAGES_DIR, consts.INDEX)

	err := utils.WriteFile(utils.GetFileNameExport(path, "ts"), []byte(content))
	return err
}

func genRoutes(rootDirName string, pages []*api.Page) error {
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

	path := fmt.Sprintf(`%s/%s/config`, rootDirName, consts.ROUTES_DIR)
	err := utils.WriteFile(utils.GetFileNameExport(path, "tsx"), []byte(content))
	return err
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
	elementString := fmt.Sprintf(`<%s {...%s["%s"]}>%s</%s>`, node.ComponentType, consts.PROPS_BY_ID, node.ID, consts.KEY_CHILDREN, node.ComponentType)

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
		reactElement.ElementString = fmt.Sprintf(`<%s {...%s["%s"]}/>`, reactElement.Component, consts.PROPS_BY_ID, reactElement.ID)
		return reactElement.ElementString
	}

	childrenString := ""
	for _, childID := range reactElement.Children {
		childrenString += mergeReactElements(childID, mapIDToReactElements)
	}

	reactElement.ElementString = strings.Replace(reactElement.ElementString, consts.KEY_CHILDREN, childrenString, 1)

	return reactElement.ElementString
}

func setUpDir(rootDirName string, pages []*api.Page) error {
	rootDir := fmt.Sprintf("%s/%s", consts.EXPORT_DIR, rootDirName)

	err := utils.CopyDirRecursively(consts.REACT_JS_BASE_DIR, rootDir)
	if err != nil {
		return err
	}

	for _, page := range pages {
		err = utils.CreateDir(fmt.Sprintf("%s/%s/%s", rootDir, consts.PAGES_DIR, page.Name))
		if err != nil {
			return err
		}
	}

	return nil
}

func getMapPagePathToPageInfo(request *api.GenReactSourceCodeRequest) map[string]*dto.PageInfo {
	mapPagePathToPageInfo := make(map[string]*dto.PageInfo)
	mapPagePathToPageName := make(map[string]string)

	for _, page := range request.Pages {
		page.Name = strcase.ToCamel(page.Name)
		mapPagePathToPageName[page.Path] = page.Name
	}

	for _, node := range request.Nodes {
		pagePath := node.PagePath
		if pagePath == consts.INVALID_PAGE_PATH {
			continue // TODO: should return err
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

		if strings.HasPrefix(node.Id, consts.ROOT_ID_PREFIX) {
			mapPagePathToPageInfo[pagePath].RootID = node.Id
		}
		mapPagePathToPageInfo[pagePath].Nodes = append(mapPagePathToPageInfo[pagePath].Nodes, &dto.Node{
			ID:            node.GetId(),
			ComponentType: node.GetType(),
			Props:         node.GetProps(),
			Children:      node.GetChildren(),
		})
	}

	return mapPagePathToPageInfo
}
