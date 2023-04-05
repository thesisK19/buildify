package utils

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func UploadFile(inputFilePath string, remoteFilePath string) (*string, error) {
	ctx := context.Background()

	bucket := "gen-code" //your bucket name
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
