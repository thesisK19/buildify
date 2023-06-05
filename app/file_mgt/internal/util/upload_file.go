package util

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"github.com/thesisK19/buildify/app/file_mgt/internal/constant"
	"google.golang.org/api/option"
)

func UploadFileWithFilePath(inputFilePath string, remoteFilePath string) (*string, error) {
	ctx := context.Background()

	bucket := constant.BUCKET //your bucket name
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("storage-key.json"))
	if err != nil {
		return nil, err
	}

	sw := storageClient.Bucket(bucket).Object(remoteFilePath).NewWriter(ctx)

	file, err := os.Open(inputFilePath) // For read access.
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(sw, file); err != nil {
		return nil, err
	}

	if err := sw.Close(); err != nil {
		return nil, err
	}

	u, _ := url.Parse("/" + bucket + "/" + sw.Attrs().Name)

	url := fmt.Sprintf("https://storage.googleapis.com%s", u.EscapedPath())
	return &url, nil
}

func UploadFile(file io.Reader, remoteFilePath string) (*string, error) {
	ctx := context.Background()

	bucket := "file-mgt" //your bucket name
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("storage-key.json"))
	if err != nil {
		return nil, err
	}

	sw := storageClient.Bucket(bucket).Object(remoteFilePath).NewWriter(ctx)

	if _, err := io.Copy(sw, file); err != nil {
		return nil, err
	}

	if err := sw.Close(); err != nil {
		return nil, err
	}

	u, _ := url.Parse("/" + bucket + "/" + sw.Attrs().Name)

	url := fmt.Sprintf("https://storage.googleapis.com%s", u.EscapedPath())
	return &url, nil
}
