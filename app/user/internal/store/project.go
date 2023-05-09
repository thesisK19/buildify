package store

import (
	"context"
	"fmt"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/user/internal/constant"
	"github.com/thesisK19/buildify/app/user/internal/dto"
	"github.com/thesisK19/buildify/app/user/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func (r *repository) CreateProject(ctx context.Context, username string, projectName string, projectType int) (*dto.Project, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateProject")
	now := time.Now().Unix()

	defaultProject, err := r.GetDefaultProjectByType(ctx, projectType)
	if err != nil {
		logger.WithError(err).Error("failed to GetDefaultProjectByType")
		return nil, err
	}

	result, err := r.db.Collection(constant.PROJECT_COLL).InsertOne(ctx, model.Project{
		Name:           projectName,
		Username:       username,
		CompressString: defaultProject.CompressString,
		CreatedAt:      now,
		UpdatedAt:      now,
	})

	if err != nil {
		logger.WithError(err).Error("failed to InsertOne")
		return nil, err
	}

	objID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		logger.WithError(err).Error("failed to convert objectID")
		return nil, err
	}

	return &dto.Project{
		Id:             objID.Hex(),
		Name:           projectName,
		CompressString: defaultProject.CompressString,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (r *repository) GetListProjects(ctx context.Context, username string) ([]*dto.Project, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetListProjects")

	var (
		result []*dto.Project
	)

	filter := bson.M{"username": username}

	// Find all Projects
	cursor, err := r.db.Collection(constant.PROJECT_COLL).Find(ctx, filter)
	if err != nil {
		logger.WithError(err).Error("failed to Find projects")
		return nil, fmt.Errorf("failed to find Projects: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var project model.Project
		if err := cursor.Decode(&project); err != nil {
			logger.WithError(err).Error("failed to decode project")
			return nil, fmt.Errorf("failed to decode project: %v", err)
		}

		result = append(result, &dto.Project{
			Id:             project.Id.Hex(),
			Name:           project.Name,
			CompressString: project.CompressString,
			CreatedAt:      project.CreatedAt,
			UpdatedAt:      project.UpdatedAt,
		})
	}

	return result, nil
}
func (r *repository) GetProject(ctx context.Context, username string, id string) (*dto.Project, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "GetProject")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID, "username": username}
	var project model.Project

	if err := r.db.Collection(constant.PROJECT_COLL).FindOne(ctx, filter).Decode(&project); err != nil {
		logger.WithError(err).Error("failed to FindOne")
		return nil, err
	}

	return &dto.Project{
		Id:             project.Id.Hex(),
		Name:           project.Name,
		CompressString: project.CompressString,
		CreatedAt:      project.CreatedAt,
		UpdatedAt:      project.UpdatedAt,
	}, nil
}

func (r *repository) UpdateProject(ctx context.Context, project model.Project) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UpdateProject")

	project.UpdatedAt = time.Now().Unix()
	
	filter := bson.M{"_id": project.Id, "username": project.Username}
	updateDoc := bson.M{"$set": project}

	result, err := r.db.Collection(constant.PROJECT_COLL).UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		logger.WithError(err).Error("failed to UpdateOne")
		return err
	}

	// Check the number of documents matched and modified
	if result.MatchedCount == 0 {
		err = fmt.Errorf("no documents matched the filter")
		logger.WithError(err).Error("failed to UpdateOne")
		return err
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("no documents were modified")
		logger.WithError(err).Error("failed to UpdateOne")
		return err
	}

	return nil
}

func (r *repository) DeleteProject(ctx context.Context, username string, id string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "DeleteProject")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID, "username": username}

	// Delete the document
	result, err := r.db.Collection(constant.PROJECT_COLL).DeleteOne(ctx, filter)
	if err != nil {
		logger.WithError(err).Error("failed to DeleteOne")
		return err
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		err = fmt.Errorf("document not found")
		logger.WithError(err).Error("failed to DeleteOne")
		return err
	}

	return nil
}
