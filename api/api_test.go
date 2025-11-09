package api

import (
	"my-project/3-struct/bins"
	"testing"
	"time"
)

var testCases = []struct {
	name     string
	expected bool
}{
	{name: "создание бина", expected: true},
	{name: "обновление бина", expected: true},
	{name: "получение бина", expected: true},
}

func TestCreateAPI(t *testing.T) {

	apis := NewAPI()
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

	err := CreateAPI(apis, arrbn1)
	if err != nil {
		t.Errorf("Не удалось создать бин %v", err)
	}

}

func TestUpdateAPI(t *testing.T) {

	apis := NewAPI()
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

	err := UpdateAPI(apis, "691096dcd0ea881f40dd911d", arrbn1)
	if err != nil {
		t.Errorf("Не удалось обновить бин %v", err)
	}

}

func TestGetAPI(t *testing.T) {

	apis := NewAPI()

	err := GetAPI(apis, "691096dcd0ea881f40dd911d")
	if err != nil {
		t.Errorf("Не удалось обновить бин %v", err)
	}

}
