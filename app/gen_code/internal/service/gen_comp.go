package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/scylladb/go-set/strset"
	"github.com/thesisK19/buildify/app/gen_code/internal/constant"
	"github.com/thesisK19/buildify/app/gen_code/internal/dto"
	"github.com/thesisK19/buildify/app/gen_code/internal/util"
)

func genComp(ctx context.Context, rootDirPath string, compName string, compInfo *dto.ComponentInfo) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "genComp")

	importComponents := strset.New()

	jsx := mergeCompNode(compInfo.RootID, 0, compInfo, importComponents)

	// create comp
	content := fmt.Sprintf(`
		import React, { FC, ReactElement } from "react"
   		import { %s } from "src/components"

		export const %s: FC<any> = (props): ReactElement => {
			const { %s } = props;
			return (
				%s
			);
		};`,
		strings.Join(importComponents.List(), ","), compName, strings.Join(compInfo.ComponentProps, ","), jsx)

	indexFilePath := fmt.Sprintf("%s/%s/%s/%s", rootDirPath, constant.USER_COMPONENT_DIR, compName, constant.INDEX_TSX)
	err := util.WriteFile(ctx, indexFilePath, []byte(content))
	if err != nil {
		logger.WithError(err).Error("failed to WriteFile")
		return err
	}

	return nil
}

func genIndexComponent(ctx context.Context, rootDirPath string, compNames []string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "genIndexComponent")

	var (
		exportComps []string
	)

	for _, comp := range compNames {
		exportComps = append(exportComps, fmt.Sprintf(`export { %s } from './%s';`, comp, comp))
	}

	content := strings.Join(exportComps, "\n")

	filePath := fmt.Sprintf(`%s/%s/%s`, rootDirPath, constant.USER_COMPONENT_DIR, constant.INDEX_TS)

	err := util.WriteFile(ctx, filePath, []byte(content))
	if err != nil {
		logger.WithError(err).Error("failed to WriteFile")
		return err
	}
	return nil
}

func findComponentProps(mapCompNameToCompInfo map[string]*dto.ComponentInfo) {
	for _, info := range mapCompNameToCompInfo {
		curIndex := 0
		props := []string{}
		recursiveCompProps(info.RootID, curIndex, &props, info)
		info.ComponentProps = props
	}
}

func recursiveCompProps(id string, curIndex int, props *[]string, info *dto.ComponentInfo) {
	comp := info.CompNodes[id]
	correspondingProp := fmt.Sprintf(`%s_%d`, comp.ComponentType, curIndex)
	*props = append(*props, correspondingProp)

	for _, childID := range comp.Children {
		curIndex += 1
		recursiveCompProps(childID, curIndex, props, info)
	}
}

func mergeCompNode(id string, curIndex int, compInfo *dto.ComponentInfo, importComponents *strset.Set) string {
	node := compInfo.CompNodes[id]
	componentType := node.ComponentType
	importComponents.Add(componentType)

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf(`<%s`, componentType))

	if node.Name != componentType {
		builder.WriteString(fmt.Sprintf(` name="%s"`, node.Name))
	}

	//  case child empty
	if len(node.Children) == 0 {
		builder.WriteString(fmt.Sprintf(` {...%s} />`, compInfo.ComponentProps[curIndex]))
		return strings.TrimSpace(builder.String())
	}

	//  case child > 0
	builder.WriteString(fmt.Sprintf(` {...%s}>`, compInfo.ComponentProps[curIndex]))

	// iteration child
	for _, childID := range node.Children {
		curIndex += 1
		builder.WriteString(mergeCompNode(childID, curIndex, compInfo, importComponents))
	}

	builder.WriteString(fmt.Sprintf(`</%s>`, node.ComponentType))
	return strings.TrimSpace(builder.String())
}
