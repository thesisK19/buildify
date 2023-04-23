package util

import (
	"archive/zip"
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	cp "github.com/otiai10/copy"
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

func RemoveFile(ctx context.Context, filename string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "RemoveFile")
	err := os.Remove(filename)

	if err != nil {
		logger.WithError(err).Error("failed to os.Remove file")
		return err
	}

	return nil
}

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

func AppendFile(ctx context.Context, filename string, data []byte) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "AppendFile")
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.WithError(err).Error("failed to os.OpenFile file")
		return err
	}

	defer file.Close()

	if _, err := file.Write(data); err != nil {
		logger.WithError(err).Error("failed to file.Write file")
		return err
	}

	return nil
}

func CopyFile(ctx context.Context, src string, dest string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CopyFile")
	bytesRead, err := os.ReadFile(src)
	if err != nil {
		logger.WithError(err).Error("failed to os.ReadFile file")
		return err
	}

	err = os.WriteFile(dest, bytesRead, 0644)
	if err != nil {
		logger.WithError(err).Error("failed to os.WriteFile file")
		return err
	}

	return nil
}

func CreateDir(ctx context.Context, path string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateDir")
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			logger.WithError(err).Error("failed to os.Mkdir")
			return err
		}
	}

	return nil
}

func CreateDirRecursively(ctx context.Context, path string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CreateDirRecursively")
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		logger.WithError(err).Error("failed to os.MkdirAll")
		return err
	}

	return nil
}

func CopyDirRecursively(ctx context.Context, src string, dest string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "CopyDirRecursively")
	// TODO: add option param outsite

	opt := cp.Options{
		Skip: func(info os.FileInfo, src, dest string) (bool, error) {
			return (strings.HasSuffix(src, ".gitkeep") || strings.HasSuffix(src, "components/.gitignore") || strings.HasSuffix(src, "components/package.json") || strings.HasSuffix(src, "components/yarn.lock")), nil
		},
	}

	err := cp.Copy(src, dest, opt)
	if err != nil {
		logger.WithError(err).Error("failed to cp.Copy")
		return err
	}

	return nil
}

func ZipDir(ctx context.Context, inputDir, outputZip string) error {
	logger := ctxlogrus.Extract(ctx).WithField("func", "ZipDir")
	// Create a new file to write the archive to
	zipFile, err := os.Create(outputZip)
	if err != nil {
		logger.WithError(err).Error("failed to os.Create")
		return err
	}
	defer zipFile.Close()

	// Create a new zip archive
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk the directory tree recursively and add files to the archive
	err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.WithError(err).Error("failed to filepath.Walk")
			return err
		}

		// Get the file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			logger.WithError(err).Error("failed to zip.FileInfoHeader")
			return err
		}

		// Set the header name to the relative path of the file
		header.Name, err = filepath.Rel(inputDir, path)
		if err != nil {
			logger.WithError(err).Error("failed to filepath.Rel")
			return err
		}

		// Check if the file is a directory
		if info.IsDir() {
			header.Name += "/"
			header.Method = zip.Store
		} else {
			// Set the compression method for the file
			header.Method = zip.Deflate
		}

		// Create a new file entry in the archive
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			logger.WithError(err).Error("failed to zipWriter.CreateHeader")
			return err
		}

		// If the file is not a directory, write its contents to the archive
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				logger.WithError(err).Error("failed to os.Open")
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				logger.WithError(err).Error("failed to io.Copy")
				return err
			}
		}

		return nil
	})
	if err != nil {
		logger.WithError(err).Error("failed to filepath.Walk")
		return err
	}

	return nil
}
