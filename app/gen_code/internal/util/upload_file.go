package util

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"google.golang.org/api/option"
)

func UploadFile(ctx context.Context, inputFilePath string, remoteFilePath string) (*string, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "UploadFile")

	bucket := "gen-code" //your bucket name TODO:
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("storage-key.json"))
	if err != nil {
		logger.WithError(err).Error("Failed to create storage.NewClient")
		return nil, err
	}

	sw := storageClient.Bucket(bucket).Object(remoteFilePath).NewWriter(ctx)

	file, err := os.Open(inputFilePath) // For read access.
	if err != nil {
		logger.WithError(err).Error("Failed to os.Open")
		return nil, err
	}

	if _, err := io.Copy(sw, file); err != nil {
		logger.WithError(err).Error("Failed to io.Copy")
		return nil, err
	}

	if err := sw.Close(); err != nil {
		logger.WithError(err).Error("Failed to sw.Close")
		return nil, err
	}

	u, _ := url.Parse("/" + bucket + "/" + sw.Attrs().Name)

	url := fmt.Sprintf("https://storage.googleapis.com%s", u.EscapedPath())
	return &url, nil
}
