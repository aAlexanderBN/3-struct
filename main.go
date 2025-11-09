package main

import (
	"flag"
	"fmt"
	"log"
	"my-project/3-struct/api"
	"my-project/3-struct/bins"
	"my-project/3-struct/storage"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./config/.env")
	//err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Затем получаем переменные
	key := os.Getenv("KEY")
	fmt.Printf("KEY: %s\n", key)

	// list := flag.Bool("list", false, "Список бинов")
	create := flag.Bool("create", false, "Создать бин")
	update := flag.Bool("update", false, "Обновить бин")
	get := flag.Bool("get", false, "Получить бин")

	// file := flag.string("file", "json.bin", "Имя файла")
	// name := flag.string("name", "", "Имя бина")
	id := flag.String("id", "691096dcd0ea881f40dd911d", "ИД бина")

	flag.Parse()

	if *create || *update {
		apis := api.NewAPI()
		arrbn1 := bins.ArrBn{
			{
				Login:     "a@a.ru",
				Password:  "HqzXKSM9o0rd",
				Url:       "http://a.ru",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Login:     "b@b.ru",
				Password:  "kCNodDqc5t0W",
				Url:       "https://b.ru",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		if *create {
			r := api.CreateAPI(apis, arrbn1)
			fmt.Println("answer create ", r)
		}

		if *update {
			r := api.UpdateAPI(apis, *id, arrbn1)
			fmt.Println("answer update ", r)
		}
	}

	if *get {
		apis := api.NewAPI()
		r := api.GetAPI(apis, *id)
		fmt.Println("answer update ", r)
	}

}

func testinterface() {

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
