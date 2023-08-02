package service

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	userApi "github.com/thesisK19/buildify/app/user/api"
	server_lib "github.com/thesisK19/buildify/library/server"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/gen_code/api"
	"github.com/thesisK19/buildify/app/gen_code/internal/constant"
	"github.com/thesisK19/buildify/app/gen_code/internal/dto"
	"github.com/thesisK19/buildify/app/gen_code/internal/util"

	"github.com/iancoleman/strcase"
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

	mapPagePathToPageInfo, mapCompNameToCompInfo, listPageName, listCompName, err := getInfoFromRequest(ctx, request)
	if err != nil {
		logger.WithError(err).Error("failed to getInfoFromRequest")
		return nil, err
	}

	findComponentProps(mapCompNameToCompInfo)

	rootDirName := strconv.FormatInt(time.Now().Unix(), constant.BASE_DECIMAL)
	rootDirPath := fmt.Sprintf("%s/%s", constant.EXPORT_DIR, rootDirName)
	outputZipPath := fmt.Sprintf("%s/%s.zip", constant.EXPORT_DIR, rootDirName)

	err = setUpDir(ctx, rootDirPath, listPageName, listCompName)
	if err != nil {
		logger.WithError(err).Error("failed to setUpDir")
		return nil, err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(mapPagePathToPageInfo)+2)

	// Generate theme
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = genTheme(ctx, rootDirPath, request.Theme)
		if err != nil {
			logger.WithError(err).Error("failed to genTheme")
			errChan <- err
		}
	}()

	// Generate database
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = s.genDatabase(ctx, rootDirPath, request.ProjectId)
		if err != nil {
			logger.WithError(err).Error("failed to genDatabase")
			errChan <- err
		}
	}()

	//
	// Generate component concurrently
	for compName, compInfo := range mapCompNameToCompInfo {
		wg.Add(1)
		go func(compName string, compInfo *dto.ComponentInfo) {
			defer wg.Done()
			err := genComp(ctx, rootDirPath, compName, compInfo)
			if err != nil {
				logger.WithError(err).Error("failed to genComp")
				errChan <- err
			}
		}(compName, compInfo)
	}

	// Generate index component concurrently
	if len(listCompName) > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = genIndexComponent(ctx, rootDirPath, listCompName)
			if err != nil {
				logger.WithError(err).Error("failed to genIndexComponent")
				errChan <- err
			}
		}()
	}

	//
	// Generate pages concurrently
	for _, pageInfo := range mapPagePathToPageInfo {
		wg.Add(1)
		go func(pageInfo *dto.PageInfo) {
			defer wg.Done()
			err := genPage(ctx, rootDirPath, pageInfo, mapCompNameToCompInfo)
			if err != nil {
				logger.WithError(err).Error("failed to genPage")
				errChan <- err
			}
		}(pageInfo)
	}

	// Generate index pages concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = genIndexPages(ctx, rootDirPath, listPageName)
		if err != nil {
			logger.WithError(err).Error("failed to genIndexPages")
			errChan <- err
		}
	}()

	// Generate routes concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = genRoutes(ctx, rootDirPath, request.Pages)
		if err != nil {
			logger.WithError(err).Error("failed to genRoutes")
			errChan <- err
		}
	}()

	// Wait for all Goroutines to finish
	wg.Wait()
	close(errChan)

	// Collect errors from errChan
	for err := range errChan {
		if err != nil {
			logger.WithError(err).Error("failed to doGenReactSourceCode")
			return nil, err
		}
	}

	// format code
	formatCode(rootDirPath, len(listCompName) > 0)

	// create zip file
	err = util.ZipDir(ctx, rootDirPath, outputZipPath)
	if err != nil {
		logger.WithError(err).Error("failed to ZipDir")
		return nil, err
	}

	// upload File
	username := server_lib.GetUsernameFromContext(ctx)
	projectName, err := s.adapters.user.InternalGetProjectBasicInfo(ctx, &userApi.InternalGetProjectBasicInfoRequest{
		Id: request.ProjectId,
	})
	if err != nil {
		logger.WithError(err).Error("failed to call user.InternalGetProjectBasicInfo")
		return nil, err
	}

	remoteFilePath := util.GenerateFileName(username, projectName.Name, constant.SOURCE_CODE+constant.ZIP_EXTENSION)

	// Remove the input directory after zipping.
	defer func() {
		go func() {
			err := os.Remove(outputZipPath)
			if err != nil {
				logger.WithError(err).Error("failed to os.Remove file")
			}
			err = os.RemoveAll(rootDirPath)
			if err != nil {
				logger.WithError(err).Error("failed to os.RemoveAll file")
			}
		}()
	}()

	url, err := util.UploadFile(ctx, outputZipPath, remoteFilePath, true, true)
	if err != nil {
		logger.WithError(err).Error("failed to UploadFile")
		return nil, err
	}

	return &api.GenReactSourceCodeResponse{
		Url: *url,
	}, nil
}

func formatCode(rootDirPath string, formatComp bool) {
	var command *exec.Cmd
	// npx prettier --write .
	if formatComp {
		command = exec.Command("prettier", "--write",
			fmt.Sprintf("%s/%s", rootDirPath, constant.PAGES_DIR),
			fmt.Sprintf("%s/%s", rootDirPath, constant.USER_COMPONENT_DIR),
			fmt.Sprintf("%s/%s", rootDirPath, constant.ROUTES_DIR),
			fmt.Sprintf("%s/%s", rootDirPath, constant.THEME_DIR),
			fmt.Sprintf("%s/%s", rootDirPath, constant.DATABASE_DIR),
		)
	} else {
		command = exec.Command("prettier", "--write",
			fmt.Sprintf("%s/%s", rootDirPath, constant.PAGES_DIR),
			fmt.Sprintf("%s/%s", rootDirPath, constant.ROUTES_DIR),
			fmt.Sprintf("%s/%s", rootDirPath, constant.THEME_DIR),
			fmt.Sprintf("%s/%s", rootDirPath, constant.DATABASE_DIR),
		)
	}
	command.Stderr = os.Stderr
	// Run the command
	command.Run()
}

func genTheme(ctx context.Context, rootDirPath string, theme string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "genTheme")

	content := fmt.Sprintf(`
		export type Theme = Record<string, any>;
		
		const MyTheme: Theme = %s

		export default MyTheme;
	`, theme)

	filePath := fmt.Sprintf(`%s/%s/%s`, rootDirPath, constant.THEME_DIR, constant.INDEX_TS)
	err := util.WriteFile(ctx, filePath, []byte(content))
	if err != nil {
		logger.WithError(err).Error("failed to WriteFile")
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
		logger.WithError(err).Error("failed to WriteFile")
		return err
	}
	return nil
}

func setUpDir(ctx context.Context, rootDirPath string, pages []string, comps []string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "setUpDir")

	err := util.CopyDirRecursively(ctx, constant.REACT_JS_BASE_DIR, rootDirPath)
	if err != nil {
		logger.WithError(err).Error("failed to CopyDirRecursively")
		return err
	}

	for _, pageName := range pages {
		err = util.CreateDir(ctx, fmt.Sprintf("%s/%s/%s", rootDirPath, constant.PAGES_DIR, pageName))
		if err != nil {
			logger.WithError(err).Error("failed to CreateDir")
			return err
		}
	}

	if len(comps) > 0 {
		err = util.CreateDir(ctx, fmt.Sprintf("%s/%s", rootDirPath, constant.USER_COMPONENT_DIR))
		if err != nil {
			logger.WithError(err).Error("failed to CreateDir")
			return err
		}
	}
	for _, compName := range comps {
		err = util.CreateDir(ctx, fmt.Sprintf("%s/%s/%s", rootDirPath, constant.USER_COMPONENT_DIR, compName))
		if err != nil {
			logger.WithError(err).Error("failed to CreateDir")
			return err
		}
	}

	return nil
}

func getInfoFromRequest(ctx context.Context, request *api.GenReactSourceCodeRequest) (map[string]*dto.PageInfo, map[string]*dto.ComponentInfo, []string, []string, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "getInfoFromRequest")

	mapPagePathToPageInfo := make(map[string]*dto.PageInfo)
	mapPagePathToPageName := make(map[string]string)
	listPageName := make([]string, 0)
	listCompName := make([]string, 0)

	mapCompNameToCompInfo := make(map[string]*dto.ComponentInfo)
	mapCompNametoCompRootID := make(map[string]string)

	for _, page := range request.GetPages() {
		page.Name = strcase.ToCamel(page.Name)
		mapPagePathToPageName[page.Path] = page.Name
		listPageName = append(listPageName, page.Name)
	}

	for _, component := range request.GetComponents() {
		name := strcase.ToCamel(component.Name)
		component.Name = name
		mapCompNametoCompRootID[name] = fmt.Sprintf(`%s_%s`, constant.ROOT_COMP_ID_PREFIX, name)
		listCompName = append(listCompName, name)
	}

	for _, node := range request.Nodes {
		pagePath := node.PagePath
		node.BelongToComponent = strcase.ToCamel(node.BelongToComponent)
		compName := node.BelongToComponent

		if pagePath == constant.INVALID_PAGE_PATH && compName == constant.INVALID_COMP_PATH {
			err := fmt.Errorf("invalid page or comp, node=%v", node)
			logger.WithError(err).Error("invalid")
			return nil, nil, nil, nil, err
		}

		// it is/belongto component
		if pagePath == constant.INVALID_PAGE_PATH {
			if _, ok := mapCompNameToCompInfo[compName]; !ok {
				mapCompNameToCompInfo[compName] = &dto.ComponentInfo{
					RootID:         mapCompNametoCompRootID[compName],
					Name:           compName,
					ComponentProps: make([]string, 0),
					CompNodes:      make(map[string]*dto.CompNode, 0),
				}
			}

			if strings.HasPrefix(node.Id, constant.ROOT_ID_PREFIX) {
				node.Id = mapCompNametoCompRootID[compName] // re-assign
			}

			mapCompNameToCompInfo[compName].CompNodes[node.Id] = &dto.CompNode{
				ID:            node.GetId(), // will not use, just use for root
				Name:          node.GetDisplayName(),
				ComponentType: node.GetType(),
				Children:      node.GetChildren(),
			}

			continue
		}

		// else if belong to page
		if _, ok := mapPagePathToPageInfo[pagePath]; !ok {
			pageName := mapPagePathToPageName[pagePath]

			mapPagePathToPageInfo[pagePath] = &dto.PageInfo{
				RootID:      "",
				Path:        pagePath,
				Name:        pageName,
				Nodes:       []*dto.Node{},
				LinkedNodes: []string{},
			}

		}

		if strings.HasPrefix(node.Id, constant.ROOT_ID_PREFIX) {
			mapPagePathToPageInfo[pagePath].RootID = node.Id
		}
		mapPagePathToPageInfo[pagePath].Nodes = append(mapPagePathToPageInfo[pagePath].Nodes, &dto.Node{
			ID:                        node.GetId(),
			Name:                      node.GetDisplayName(),
			ComponentType:             node.GetType(),
			BelongToUserComponentType: compName,
			// CorrespondingProp calc later
			Props:    node.GetProps(),
			Children: node.GetChildren(),
		})
		mapPagePathToPageInfo[pagePath].LinkedNodes = append(mapPagePathToPageInfo[pagePath].LinkedNodes, node.LinkedNodes...)
	}

	return mapPagePathToPageInfo, mapCompNameToCompInfo, listPageName, listCompName, nil
}
