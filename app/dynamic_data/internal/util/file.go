package util

import (
	"context"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
)

// https://zetcode.com/golang/file/
// https://freshman.tech/snippets/go/create-directory-if-not-exist/

// If the file already exists, it is truncated. If the file does not exist, it is created with mode 0666
func CreateFile(ctx context.Context, filename string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateFile")

	file, err := os.Create(filename)
	if err != nil {
		logger.WithError(err).Error("failed to os.Create file")
		return err
	}

	defer file.Close()

	return nil
}

// func RemoveFile(ctx context.Context, filename string) error {
// 	logger := ctxlogrus.Extract(ctx).WithField("func", "RemoveFile")
// 	err := os.Remove(filename)

// 	if err != nil {
// 		logger.WithError(err).Error("failed to os.Remove file")
// 		return err
// 	}

// 	return nil
// }

// func RemoveAllFiles(ctx context.Context, folderPath string) error {
// 	logger := ctxlogrus.Extract(ctx).WithField("func", "RemoveAllFiles")
// 	err := os.RemoveAll(folderPath)

// 	if err != nil {
// 		logger.WithError(err).Error("failed to os.RemoveAll files")
// 		return err
// 	}

// 	return nil
// }

func ReadFile(ctx context.Context, filename string) ([]byte, error) {
	logger := ctxlogrus.Extract(ctx).WithField("func", "ReadFile")

	content, err := os.ReadFile(filename)
	if err != nil {
		logger.WithError(err).Error("failed to os.ReadFile file")
		return nil, err
	}

	return content, nil
}

func WriteFile(ctx context.Context, filename string, data []byte) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "WriteFile")

	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		logger.WithError(err).Error("failed to os.WriteFile file")
		return err
	}

	return nil
}
