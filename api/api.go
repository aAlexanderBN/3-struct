package api

import (
	"fmt"
	"my-project/3-struct/config"
)

type ApiDb struct {
	connect string
	key     string
}

func (db *ApiDb) Write(b string) error {

	fmt.Println("FileDb write ", db.connect, b)
	return nil
}

func (db *ApiDb) Read() (string, error) {

	var b string
	return b, nil
}

func NewAPI() *ApiDb {

	return &ApiDb{
		connect: "",
		key:     config.New().Key,
	}
}
