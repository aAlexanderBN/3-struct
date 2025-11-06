package bins

import (
	"os"
	"time"
)

type BinDb struct {
	FileName string
}

type BinList []*Bin

type Bin struct {
	Id       string
	Private  bool
	CreateAt time.Time
	Name     string
}

func NewBin(id string, private bool, createAt time.Time, name string) *Bin {

	return &Bin{
		Id:       id,
		Private:  private,
		CreateAt: createAt,
		Name:     name,
	}
}

func (db *BinDb) Write(b string) error {

	file, err := os.Create(db.FileName)
	if err != nil {
		return err
	}
	// var bArr []Bin

	defer file.Close()
	// data, err := json.Marshal(bArr)
	// if err != nil {
	// 	return err
	// }
	_, err = file.Write([]byte(b))
	if err != nil {
		return err
	}

	return nil
}

func (db *BinDb) Read() (string, error) {

	//var b []Bin

	data, err := os.ReadFile(db.FileName)
	if err != nil {
		return "", err
	}
	// err = json.Unmarshal(data, &b)
	// if err != nil {
	// 	return b, err
	// }
	return string(data), err
}
