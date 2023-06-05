package util

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/thesisK19/buildify/app/dynamic_data/internal/constant"
	"google.golang.org/api/option"
)

func GenerateFileName(username string, projectName string, additionalInfo string) string {
	// Concatenate username and projectName
	fileName := username + "_" + projectName

	// Add additionalInfo if it is not empty
	if additionalInfo != "" {
		fileName += "_" + additionalInfo
	}

	// Replace spaces with underscores
	fileName = strings.ReplaceAll(fileName, " ", "_")
	return fileName
}

func UploadFile(ctx context.Context, inputFilePath string, remoteFilePath string, deleteAfterDuration bool, noCache bool) (*string, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UploadFile")

	bucket := constant.BUCKET //your bucket name TODO:
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("storage-key.json"))
	if err != nil {
		logger.WithError(err).Error("failed to create storage.NewClient")
		return nil, err
	}

	// Create a new writer
	sw := storageClient.Bucket(bucket).Object(remoteFilePath).NewWriter(ctx)
	// Set the Cache-Control metadata
	if noCache {
		sw.ObjectAttrs.CacheControl = "max-age=0, no-cache"
	}

	file, err := os.Open(inputFilePath) // For read access.
	if err != nil {
		logger.WithError(err).Error("failed to os.Open")
		return nil, err
	}

	if _, err := io.Copy(sw, file); err != nil {
		logger.WithError(err).Error("failed to io.Copy")
		return nil, err
	}

	if err := sw.Close(); err != nil {
		logger.WithError(err).Error("failed to sw.Close")
		return nil, err
	}

	u, _ := url.Parse("/" + bucket + "/" + sw.Attrs().Name)

	url := fmt.Sprintf("https://storage.googleapis.com%s", u.EscapedPath())

	if deleteAfterDuration {
		// Start a goroutine to delete the file after the specified duration
		go func() {
			// Sleep for the specified duration
			time.Sleep(constant.DETELE_REMOTE_FILE_DURATION)

			obj := storageClient.Bucket(bucket).Object(remoteFilePath)
			if err := obj.Delete(context.Background()); err != nil {
				logger.WithError(err).Error("Failed to delete existing file" + remoteFilePath)
			}
		}()
	}

	return &url, nil
}
