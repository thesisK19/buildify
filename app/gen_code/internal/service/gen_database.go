package service

import (
	"context"
	"encoding/json"
	"fmt"

	dynamicDataApi "github.com/thesisK19/buildify/app/dynamic_data/api"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/gen_code/internal/constant"
	"github.com/thesisK19/buildify/app/gen_code/internal/util"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *Service) genDatabase(ctx context.Context, rootDirPath string, projectId string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "genDatabase")

	// call dynamic data service to get database
	resp, err := s.adapters.dynamicData.GetCollectionMapping(ctx, &dynamicDataApi.GetCollectionMappingRequest{
		ProjectId: projectId,
	})
	if err != nil {
		logger.WithError(err).Error("failed to GetCollectionMapping from dynamic data service")
		return err
	}
	jsonString, err := protojson.Marshal(resp)
	if err != nil {
		logger.WithError(err).Error("failed to Marshal message proto")
		return err
	}

	jsonDataValue, err := getJsonDataValueOfDatabase(string(jsonString))
	if err != nil {
		logger.WithError(err).Error("failed to getJsonDataValueOfDatabase")
		return err
	}

	content := fmt.Sprintf(
		`type Id = number;
		
		type Collection = {
			name: string;
			dataKeys: string[];
			dataTypes: string[];
			documents: Record<Id, any>;
		};
		
		export type Database = Record<Id, Collection>;
		
		export const MyDatabase: Database = %s
	`, jsonDataValue)

	filePath := fmt.Sprintf(`%s/%s/%s`, rootDirPath, constant.DATABASE_DIR, constant.INDEX_TS)
	err = util.WriteFile(ctx, filePath, []byte(content))
	if err != nil {
		logger.WithError(err).Error("failed to WriteFile")
		return err
	}
	return nil
}

func getJsonDataValueOfDatabase(jsonString string) (string, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return "", err
	}

	dataValue, ok := data["data"]
	if !ok {
		return "{}", nil
	}

	dataJSON, err := json.Marshal(dataValue)
	if err != nil {
		return "", err
	}

	return string(dataJSON), nil
}
