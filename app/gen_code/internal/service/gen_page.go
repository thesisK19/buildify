package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/scylladb/go-set/strset"
	"github.com/thesisK19/buildify/app/gen_code/internal/constant"
	"github.com/thesisK19/buildify/app/gen_code/internal/dto"
	"github.com/thesisK19/buildify/app/gen_code/internal/util"
	"golang.org/x/exp/slices"
)

func genPage(ctx context.Context, rootDirPath string, pageInfo *dto.PageInfo, mapCompNameToCompInfo map[string]*dto.ComponentInfo) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "genPage")

	importComponents, importUserComponents, mapIDToReactElements, propsByIds := getReactElementInfoFromNodes(pageInfo.Nodes, pageInfo.LinkedNodes)

	reactElementString := mergeReactElements(pageInfo.RootID, mapIDToReactElements, mapCompNameToCompInfo)

	// create props
	propsString := fmt.Sprintf("export const %s = {%s}", constant.REF_PROPS, strings.Join(propsByIds, ","))

	propsFilePath := fmt.Sprintf("%s/%s/%s/%s", rootDirPath, constant.PAGES_DIR, pageInfo.Name, constant.PROPS_TSX)
	err := util.WriteFile(ctx, propsFilePath, []byte(propsString))
	if err != nil {
		logger.WithError(err).Error("failed to WriteFile")
		return err
	}

	// importUserComponents
	importUserComponentsString := ""
	if len(importUserComponents) > 0 {
		joinString := strings.Join(importUserComponents, ",")
		importUserComponentsString = "\n" + fmt.Sprintf(`import { %s } from "src/user-components";`, joinString)
	}
	// create page
	content := fmt.Sprintf(`
		import React, { FC, ReactElement } from "react"
   		import { %s } from "src/components"%s
    	import { %s } from "./props"
		import { useGetValuesFromReferencedProps } from "src/hooks/useGetValuesFromReferencedProps";

    	const %s: FC = (): ReactElement => {
			const props = useGetValuesFromReferencedProps(refProps);
			return (
				%s
			)
		}
		
		export default %s`,
		strings.Join(importComponents, ","), importUserComponentsString, constant.REF_PROPS, pageInfo.Name, reactElementString, pageInfo.Name)

	indexFilePath := fmt.Sprintf("%s/%s/%s/%s", rootDirPath, constant.PAGES_DIR, pageInfo.Name, constant.INDEX_TSX)
	err = util.WriteFile(ctx, indexFilePath, []byte(content))
	if err != nil {
		logger.WithError(err).Error("failed to WriteFile")
		return err
	}

	return nil
}

func genIndexPages(ctx context.Context, rootDirPath string, pageNames []string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "genIndexPages")

	var (
		exportPages []string
		importPages []string
	)

	for _, pageName := range pageNames {
		importPages = append(importPages, fmt.Sprintf(`import %s from "./%s"`, pageName, pageName))
		exportPages = append(exportPages, pageName)
	}

	content := fmt.Sprintf("%s \n\n export {%s}", strings.Join(importPages, "\n"), strings.Join(exportPages, ","))

	filePath := fmt.Sprintf(`%s/%s/%s`, rootDirPath, constant.PAGES_DIR, constant.INDEX_TS)

	err := util.WriteFile(ctx, filePath, []byte(content))
	if err != nil {
		logger.WithError(err).Error("failed to WriteFile")
		return err
	}
	return nil
}

func getReactElementInfoFromNodes(nodes []*dto.Node, linkedNodes []string) ([]string, []string, map[string]*dto.ReactElement, []string) {
	components := strset.New()
	userComponents := strset.New()
	mapIDToReactElements := map[string]*dto.ReactElement{}
	propsByIds := []string{}

	for _, node := range nodes {
		if slices.Contains(linkedNodes, node.ID) {
			continue
		}

		if node.BelongToUserComponentType != "" {
			userComponents.Add(node.BelongToUserComponentType)
		} else {
			components.Add(node.ComponentType)
		}

		reactElement := genReactElementFromNode(node)
		mapIDToReactElements[node.ID] = reactElement

		propsByIds = append(propsByIds, fmt.Sprintf(`"%s": %s`, reactElement.ID, reactElement.Props))
	}

	// sort prop by id
	sort.Slice(propsByIds, func(i, j int) bool {
		return propsByIds[i] < propsByIds[j]
	})

	return components.List(), userComponents.List(), mapIDToReactElements, propsByIds
}

func genReactElementFromNode(node *dto.Node) *dto.ReactElement {
	hasChildren := len(node.Children) > 0
	elementString := ""

	// only handle case not belong to user component
	if node.BelongToUserComponentType == "" {
		elementString = getElementString(*node, hasChildren)
	}

	return &dto.ReactElement{
		ID:                        node.ID,
		Props:                     node.Props,
		Component:                 node.ComponentType,
		ElementString:             elementString,
		Children:                  node.Children,
		BelongToUserComponentType: node.BelongToUserComponentType,
	}
}
func mergeReactElements(ID string, mapIDToReactElements map[string]*dto.ReactElement, mapCompNameToCompInfo map[string]*dto.ComponentInfo) string {
	reactElement, ok := mapIDToReactElements[ID]
	if !ok {
		return ""
	}

	// handle for user-component, the first one is root user component

	// <Product
	// Container_1={props.Container_kikYs}
	// Image_1={props.Image_ECSlM}
	// Text_1={props.Text_qJwWg}
	// Text_2={props.Text_SjIgw}
	// Text_3={props.Text_tAudj}
	// Button_1={props.Button_tSQYp} />
	if reactElement.BelongToUserComponentType != "" {
		elementString := fmt.Sprintf(`<%s %s />`, reactElement.BelongToUserComponentType, constant.KEY_CHILDREN_PROPS)
		props := []string{}
		componentInfo := mapCompNameToCompInfo[reactElement.BelongToUserComponentType]
		if componentInfo == nil {
			return ""
		}

		curIndex := constant.START_CUR_INDEX
		recursiveCorrespondingProps(ID, &curIndex, &props, componentInfo.ComponentProps, mapIDToReactElements)
		reactElement.ElementString = strings.Replace(elementString, constant.KEY_CHILDREN_PROPS, strings.Join(props, "\n"), 1)
		return reactElement.ElementString
	}

	if len(reactElement.Children) == 0 {
		return reactElement.ElementString
	}

	childrenString := ""
	for _, childID := range reactElement.Children {
		childrenString += mergeReactElements(childID, mapIDToReactElements, mapCompNameToCompInfo)
	}

	reactElement.ElementString = strings.Replace(reactElement.ElementString, constant.KEY_CHILDREN, childrenString, 1)

	return reactElement.ElementString
}

func parseProps(jsonStr string) (dto.ImportantProps, error) {
	var props dto.ImportantProps
	err := json.Unmarshal([]byte(jsonStr), &props)
	return props, err
}

func getElementString(node dto.Node, hasChildren bool) string {
	id := node.ID
	componentType := node.ComponentType
	name := node.Name

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf(`<%s`, componentType))

	if name != componentType {
		builder.WriteString(fmt.Sprintf(` name="%s"`, name))
	}

	importantProps, err := parseProps(node.Props)
	if err == nil && importantProps.Text != "" {
		builder.WriteString(fmt.Sprintf(` text="%s"`, importantProps.Text))
	}

	if hasChildren {
		builder.WriteString(fmt.Sprintf(` {...%s.%s}>%s</%s>`, constant.PROPS, id, constant.KEY_CHILDREN, node.ComponentType))
	} else {
		builder.WriteString(fmt.Sprintf(` {...%s.%s}/>`, constant.PROPS, id))
	}

	return strings.TrimSpace(builder.String())
}

// Container_1={props.Container_kikYs}
func recursiveCorrespondingProps(id string, curIndex *int, resultProps *[]string, componentProps []string, mapIDToReactElements map[string]*dto.ReactElement) {
	reactElement := mapIDToReactElements[id] // sure not get err

	*resultProps = append(*resultProps, fmt.Sprintf(`%s={props.%s}`, componentProps[*curIndex], id))

	for _, childID := range reactElement.Children {
		*curIndex += 1
		recursiveCorrespondingProps(childID, curIndex, resultProps, componentProps, mapIDToReactElements)
	}
}
