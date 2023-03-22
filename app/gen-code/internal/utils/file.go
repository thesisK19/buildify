package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"thesis/be/app/gen-code/internal/consts"

	cp "github.com/otiai10/copy"
)

// https://zetcode.com/golang/file/
// https://freshman.tech/snippets/go/create-directory-if-not-exist/

// If the file already exists, it is truncated. If the file does not exist, it is created with mode 0666
func CreateFile(filename string) error {
	file, err := os.Create("filename")
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
			return strings.HasSuffix(src, ".git"), nil
		},
	}

	err := cp.Copy(src, dest, opt)
	if err != nil {
		log.Println(err) // TODO:
		return err
	}

	return nil
}

func GetFileNameExport(name string, extention string) string {
	return fmt.Sprintf(`%s/%s.%s`, consts.EXPORT_DIR, name, extention)
}
