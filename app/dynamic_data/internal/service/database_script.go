package service

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	userApi "github.com/thesisK19/buildify/app/user/api"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/api"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/constant"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/dto"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/util"
	errors_lib "github.com/thesisK19/buildify/library/errors"
	server_lib "github.com/thesisK19/buildify/library/server"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) GetDatabaseScript(ctx context.Context, in *api.GetDatabaseScriptRequest) (*api.GetDatabaseScriptResponse, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetListCollections")

	username := server_lib.GetUsernameFromContext(ctx)

	var (
		projectObjectId = primitive.NilObjectID
		err             error
	)

	projectObjectId, err = primitive.ObjectIDFromHex(in.ProjectId)
	if err != nil {
		logger.WithError(err).Error("failed to convert ObjectIDFromHex")
		return nil, errors_lib.ToInvalidArgumentError(err)
	}

	database, err := s.repository.GetListCollections(ctx, username, projectObjectId)
	if err != nil {
		logger.WithError(err).Error("failed to repo.GetListCollections")
		return nil, err
	}

	script := generateSQLScript(database, in.DatabaseSystem)

	filename := strconv.FormatInt(time.Now().Unix(), constant.BASE_DECIMAL)
	filePath := fmt.Sprintf("%s/%s.sql", constant.EXPORT_DIR, filename)

	err = util.WriteFile(ctx, filePath, []byte(script))
	if err != nil {
		logger.WithError(err).Error("failed to WriteFile")
		return nil, err
	}

	// upload remote
	projectName, err := s.adapters.user.InternalGetProjectBasicInfo(ctx, &userApi.InternalGetProjectBasicInfoRequest{
		Id: in.ProjectId,
	})
	if err != nil {
		logger.WithError(err).Error("failed to call user.InternalGetProjectBasicInfo")
		return nil, err
	}

	remoteFilePath := util.GenerateFileName(username, projectName.Name, in.DatabaseSystem.String()+constant.SQL_EXTENSION)

	url, err := util.UploadFile(ctx, filePath, remoteFilePath, true, true)
	if err != nil {
		logger.WithError(err).Error("failed to UploadFile")
		return nil, err
	}

	go func() {
		err := os.Remove(filePath)
		if err != nil {
			logger.WithError(err).Error("failed to os.Remove file")
		}
	}()

	return &api.GetDatabaseScriptResponse{
		Url: *url,
	}, nil
}

func generateSQLScript(database *dto.ListCollections, databaseSystem api.DatabaseSystem) string {
	var script strings.Builder

	for _, collection := range database.Collections {
		tableName := formatTableName(collection.Name, databaseSystem)

		createTableStatement := generateCreateTableStatement(tableName, collection.DataKeys, collection.DataTypes, databaseSystem)
		insertStatements := generateInsertStatements(tableName, collection, database.Documents, databaseSystem)

		script.WriteString(createTableStatement)
		script.WriteString(insertStatements)
	}

	return script.String()
}

// Helper function to generate the CREATE TABLE statement
func generateCreateTableStatement(tableName string, dataKeys []string, dataTypes []int32, databaseSystem api.DatabaseSystem) string {
	var createTableStatement strings.Builder

	createTableStatement.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tableName))
	createTableStatement.WriteString(generatePrimaryKeySyntax(databaseSystem))
	createTableStatement.WriteString(", ")

	for i, dataKey := range dataKeys {
		dataType := convertDataType(constant.DataType(dataTypes[i]), databaseSystem)
		createTableStatement.WriteString(fmt.Sprintf("%s %s", dataKey, dataType))

		if i < len(dataKeys)-1 {
			createTableStatement.WriteString(", ")
		}
	}

	createTableStatement.WriteString(");\n\n")

	return createTableStatement.String()
}

// Helper function to generate the auto-incrementing primary key syntax
func generatePrimaryKeySyntax(databaseSystem api.DatabaseSystem) string {
	switch databaseSystem {
	case api.DatabaseSystem_MYSQL:
		return "id INT AUTO_INCREMENT PRIMARY KEY"
	case api.DatabaseSystem_POSTGRES:
		return "id SERIAL PRIMARY KEY"
	case api.DatabaseSystem_SQLSERVER:
		return "id INT IDENTITY(1,1) PRIMARY KEY"
	case api.DatabaseSystem_SQLITE:
		return "id INTEGER PRIMARY KEY AUTOINCREMENT"
	default:
		return "id INT AUTO_INCREMENT PRIMARY KEY"
	}
}

// Helper function to generate the INSERT statements
func generateInsertStatements(tableName string, collection dto.Collection, documents []dto.Document, databaseSystem api.DatabaseSystem) string {
	var insertStatements strings.Builder

	insertStatements.WriteString(fmt.Sprintf("INSERT INTO %s (", tableName))
	for i, dataKey := range collection.DataKeys {
		insertStatements.WriteString(dataKey)

		if i < len(collection.DataKeys)-1 {
			insertStatements.WriteString(", ")
		}
	}

	insertStatements.WriteString(") VALUES\n")

	values := make([]string, 0)
	for _, document := range documents {
		if document.CollectionId == collection.Id {
			valueSet := generateValues(collection, document, databaseSystem)
			values = append(values, valueSet)
		}
	}

	insertStatements.WriteString(strings.Join(values, ",\n"))
	insertStatements.WriteString(";\n\n")

	return insertStatements.String()
}

// Helper function to generate the values for INSERT statements
func generateValues(collection dto.Collection, document dto.Document, databaseSystem api.DatabaseSystem) string {
	var values strings.Builder

	values.WriteString("(")

	for i, dataKey := range collection.DataKeys {
		values.WriteString(formatValue(document.Data[dataKey], databaseSystem))

		if i < len(collection.DataKeys)-1 {
			values.WriteString(", ")
		}
	}

	values.WriteString(")")

	return values.String()
}

// Helper function to format table name based on the database system
func formatTableName(name string, databaseSystem api.DatabaseSystem) string {
	// Remove leading and trailing spaces
	name = strings.TrimSpace(name)

	// Remove consecutive spaces
	re := regexp.MustCompile(`\s+`)
	name = re.ReplaceAllString(name, " ")

	// Replace spaces with underscores
	name = strings.ReplaceAll(name, " ", "_")

	switch databaseSystem {
	case api.DatabaseSystem_MYSQL, api.DatabaseSystem_SQLITE:
		return "`" + strings.ReplaceAll(name, " ", "_") + "`"
	case api.DatabaseSystem_POSTGRES:
		return `"` + strings.ReplaceAll(name, " ", "_") + `"`
	case api.DatabaseSystem_SQLSERVER:
		return "[" + strings.ReplaceAll(name, " ", "_") + "]"
	default:
		return name
	}
}

// Helper function to convert data type based on the database system
func convertDataType(dataType constant.DataType, databaseSystem api.DatabaseSystem) string {
	switch databaseSystem {
	case api.DatabaseSystem_MYSQL:
		switch dataType {
		case constant.STRING:
			return "VARCHAR(255)"
		case constant.NUMBER:
			return "INT"
		default:
			return "VARCHAR(255)"
		}
	case api.DatabaseSystem_POSTGRES:
		switch dataType {
		case constant.STRING:
			return "VARCHAR(255)"
		case constant.NUMBER:
			return "INTEGER"
		default:
			return "VARCHAR(255)"
		}
	case api.DatabaseSystem_SQLSERVER:
		switch dataType {
		case constant.STRING:
			return "NVARCHAR(255)"
		case constant.NUMBER:
			return "INT"
		default:
			return "NVARCHAR(255)"
		}
	case api.DatabaseSystem_SQLITE:
		switch dataType {
		case constant.STRING:
			return "TEXT"
		case constant.NUMBER:
			return "INTEGER"
		default:
			return "TEXT"
		}
	default:
		return "VARCHAR(255)"
	}
}

// Helper function to format values for SQL statements based on the database system
func formatValue(value interface{}, databaseSystem api.DatabaseSystem) string {
	switch databaseSystem {
	case api.DatabaseSystem_MYSQL, api.DatabaseSystem_POSTGRES, api.DatabaseSystem_SQLSERVER, api.DatabaseSystem_SQLITE:
		return fmt.Sprintf("'%v'", value)
	default:
		return fmt.Sprintf("'%v'", value)
	}
}
