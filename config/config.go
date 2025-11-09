package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func New() *Config {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("❌ Не удалось загрузить .env файл:", err)
	}
	return &Config{
		Key: os.Getenv("KEY"),
	}

}
