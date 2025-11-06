package bins

import (
	"time"
)

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
