package api

import "fmt"

type ApiDb struct {
	connect string
}

func (db *ApiDb) Write(b string) error {

	fmt.Println("FileDb write ", db.connect, b)
	return nil
}

func (db *ApiDb) Read() (string, error) {

	var b string
	return b, nil
}
