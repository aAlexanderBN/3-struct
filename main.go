package main

import (
	"fmt"
	"my-project/3-struct/bins"
	"my-project/3-struct/storage"
	"time"
)

func main() {

	b1 := bins.NewBin("1", true, time.Now(), "test")
	b2 := bins.NewBin("2", false, time.Now(), "test2")
	var blist bins.BinList
	blist = append(blist, b1, b2)

	fmt.Println("Тест1 ", *blist[0], *blist[1])

	var barr []bins.Bin
	barr = append(barr, *b1, *b2)
	err := storage.SaveBin("test.json", barr)
	if err != nil {
		fmt.Println("Ошибка сохранения")
		return
	}
	data2, err2 := storage.ReadBin("test.json")
	if err2 != nil {
		fmt.Println("Ошибка сохранения")
		return
	}

	fmt.Println("тест2б запись и чтение в файл ", data2[0], data2[1])
}
