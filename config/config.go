package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var APIKey string

func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var isBool bool
	APIKey, isBool = os.LookupEnv("APIKEY")
	if isBool == false {
		log.Printf("API Key não encontrada")
	}

}

const BaseURL = "https://api.brawlstars.com/v1"
const Port = "8080"
const Tag = "%23V0CJ2J"
