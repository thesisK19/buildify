package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"thesis/be/app/gen-code/api"
	"thesis/be/app/gen-code/internal/consts"
	"thesis/be/app/gen-code/internal/dto"
	"thesis/be/app/gen-code/internal/utils"
	"time"

	"github.com/scylladb/go-set/strset"
)

func (s *Service) GenReactSourceCode(ctx context.Context, request *api.GenReactSourceCodeRequest) (*api.GenReactSourceCodeResponse, error) {
	return s.doGenReactSourceCode(ctx, request)
}

func (s *Service) doGenReactSourceCode(ctx context.Context, request *api.GenReactSourceCodeRequest) (*api.GenReactSourceCodeResponse, error) {
	mapPagePathToPageInfo := getMapPagePathToPageInfo(request)

	rootDirName := strconv.FormatInt(time.Now().Unix(), consts.BASE_DECIMAL)
	err := setUpDir(rootDirName, request.Pages)

	// for pagePath, pageInfo := range mapPagePathToPageInfo {
	// 	genPage(rootDirName, pageInfo)
	// }

	// genIndexPages(rootDirName, request.Pages)

	// genRoutes(rootDirName, request.Pages)

	return nil, err
}

func genPage(rootDirName string, pageInfo dto.PageInfo) {
	components, mapIDToReactElements, propsByIds := getReactElementInfoFromNodes(pageInfo.Nodes)

	mergeReactElements(pageInfo.RootID, mapIDToReactElements)

	// create props
}

func getReactElementInfoFromNodes(nodes []*dto.Node) ([]string, map[string]*dto.ReactElement, []string) {
	components := strset.New()
	mapIDToReactElements := map[string]*dto.ReactElement{}
	propsByIds := []string{}

	for _, node := range nodes {
		reactElement := genReactElementFromNode(node)
		mapIDToReactElements[node.ID] = reactElement
		components.Add(reactElement.Component)
		propsByIds = append(propsByIds, fmt.Sprintf(`"%s": %s`), reactElement.ID, reactElement.Props)
	}

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
		childrenString = mergeReactElements(childID, mapIDToReactElements)
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
		mapPagePathToPageName[page.Path] = page.Name
	}

	for _, node := range request.Nodes {
		pagePath := node.PagePath
		if pagePath == consts.INVALID_PAGE_PATH {
			continue
		}

		_, ok := mapPagePathToPageInfo[pagePath]
		if !ok {
			pageName, _ := mapPagePathToPageName[pagePath]

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
