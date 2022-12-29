package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvirontment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env")
	}
}
