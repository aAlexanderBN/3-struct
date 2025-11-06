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
	db := bins.BinDb{
		FileName: "test.json",
	}

	_ = db

	bintest := &bins.BinDb{
		FileName: "interfacetest.json",
	}

	var testinterface storage.DataStorage
	testinterface = bintest

	data := fmt.Sprintf("%v", barr)

	erri := testinterface.Write(data)

	if erri != nil {
		fmt.Println("Ошибка записи интерфейса")
		return
	}

	data, erri = testinterface.Read()
	if erri != nil {
		fmt.Println("Ошибка чтения интерфейса")
		return
	}

	fmt.Println("Прочитали файл ", data)
}
