package file

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileDb struct {
	FileName string
}

func (db *FileDb) Write(b string) error {

	fmt.Println("FileDb write ", db.FileName, b)
	return nil
}

func (db *FileDb) Read() (string, error) {

	var b string
	return b, nil
}

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
