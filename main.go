package main

import (
	"fmt"
	"time"
)

type BinList []*Bin

type Bin struct {
	Id       string
	Private  bool
	CreateAt time.Time
	Name     string
}

func main() {

	b1 := NewBin("1", true, time.Now(), "test")
	b2 := NewBin("2", false, time.Now(), "test2")
	var blist BinList
	blist = append(blist, b1, b2)
	fmt.Println(*blist[0], *blist[1])

}

func NewBin(id string, private bool, createAt time.Time, name string) *Bin {

	return &Bin{
		Id:       id,
		Private:  private,
		CreateAt: createAt,
		Name:     name,
	}
}
