package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"my-project/3-struct/bins"
	"my-project/3-struct/config"
	"net/http"
	"time"
)

type ApiDb struct {
	connect string
	key     string
}

func GetAPI(apis *ApiDb, id string) error {

	// Тестовый запрос к API пакетов
	url := "https://api.jsonbin.io/v3/b/" + id

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Master-Key", apis.key)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Response: %s\n", string(body))

	return nil
}

func UpdateAPI(apis *ApiDb, id string, arrbn bins.ArrBn) error {

	jsonData, err := json.Marshal(arrbn)
	if err != nil {

		return errors.New("Err_Marshal")
	}

	req, err := http.NewRequest("PUT", "https://api.jsonbin.io/v3/b/"+id, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Ошибка создания запроса: %v\n", err)
		return nil
	}

	// Устанавливаем заголовки
	req.Header.Set("X-Master-Key", apis.key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "MyApp/1.0")

	// Создаем HTTP клиент с таймаутом
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка при выполнении запроса: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка чтения ответа: %v\n", err)
		return nil
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(body))

	return nil
}

func CreateAPI(apis *ApiDb, arrbn bins.ArrBn) error {

	jsonData, err := json.Marshal(arrbn)
	if err != nil {

		return errors.New("Err_Marshal")
	}

	req, err := http.NewRequest("POST", "https://api.jsonbin.io/v3/b", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Ошибка создания запроса: %v\n", err)
		return nil
	}

	// Устанавливаем заголовки
	req.Header.Set("X-Master-Key", apis.key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "MyApp/1.0")

	// Создаем HTTP клиент с таймаутом
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Выполняем запрос
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка при выполнении запроса: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка чтения ответа: %v\n", err)
		return nil
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", string(body))

	return nil
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
