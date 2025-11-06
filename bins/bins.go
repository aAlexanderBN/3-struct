package bins

import (
	"encoding/json"
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

func (db *BinDb) Write(b []Bin) error {

	file, err := os.Create(db.FileName)
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

func (db *BinDb) Read() ([]Bin, error) {

	var b []Bin

	data, err := os.ReadFile(db.FileName)
	if err != nil {
		return b, err
	}
	err = json.Unmarshal(data, &b)
	if err != nil {
		return b, err
	}
	return b, err
}
