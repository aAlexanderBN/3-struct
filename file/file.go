package file

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

//"os"

func ReadFile(filename string) ([]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func IsJSON(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}
