package main

import (
	"fmt"
	"my-project/3-struct/bins"
	"time"
)

func main() {

	b1 := bins.NewBin("1", true, time.Now(), "test")
	b2 := bins.NewBin("2", false, time.Now(), "test2")
	var blist bins.BinList
	blist = append(blist, b1, b2)
	fmt.Println(*blist[0], *blist[1])

}
