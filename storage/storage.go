package storage

import (
	"encoding/json"
	"my-project/3-struct/bins"
	"os"
)

func SaveBin(filename string, b []bins.Bin) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadBin(filename string) ([]bins.Bin, error) {

	var b []bins.Bin

	data, err := os.ReadFile(filename)
	if err != nil {
		return b, err
	}
	err = json.Unmarshal(data, &b)
	if err != nil {
		return b, err
	}
	return b, err
}
