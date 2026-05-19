package static

import (
	"errors"
	"os"
	"path"
)

func GetStaticFiles(url string) ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := path.Join(cwd, "public", url)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("File not found")
	}
	return data, nil
}

func TryFiles(url string) ([]byte, error) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := path.Join(cwd, "public", url)
	_, err = os.Stat(path)
	if err == nil {
		return GetStaticFiles(url)
	}

	return nil, errors.New("File not found")
}
