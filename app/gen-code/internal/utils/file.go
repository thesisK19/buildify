package utils

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
)

// https://zetcode.com/golang/file/
// https://freshman.tech/snippets/go/create-directory-if-not-exist/

// If the file already exists, it is truncated. If the file does not exist, it is created with mode 0666
func CreateFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err) // TODO:
		return err
	}

	defer file.Close()

	return nil
}

func RemoveFile(filename string) error {
	err := os.Remove(filename)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func ReadFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return content, nil
}

func WriteFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func AppendFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer file.Close()

	if _, err := file.Write(data); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func CopyFile(src string, dest string) error {
	bytesRead, err := os.ReadFile(src)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = os.WriteFile(dest, bytesRead, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func CreateDir(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err) // TODO:
			return err
		}
	}

	return nil
}

func CreateDirRecursively(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err) // TODO:
		return err
	}

	return nil
}

func CopyDirRecursively(src string, dest string) error {
	// TODO: add option param

	opt := cp.Options{
		Skip: func(info os.FileInfo, src, dest string) (bool, error) {
			return (strings.HasSuffix(src, "components/.gitignore") || strings.HasSuffix(src, "components/package.json") || strings.HasSuffix(src, "components/yarn.lock")), nil
		},
	}

	err := cp.Copy(src, dest, opt)
	if err != nil {
		log.Println(err) // TODO:
		return err
	}

	return nil
}

func ZipDir(inputDir, outputZip string) error {
	// Create a new file to write the archive to
	zipFile, err := os.Create(outputZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create a new zip archive
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk the directory tree recursively and add files to the archive
	err = filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set the header name to the relative path of the file
		header.Name, err = filepath.Rel(inputDir, path)
		if err != nil {
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
			return err
		}

		// If the file is not a directory, write its contents to the archive
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
