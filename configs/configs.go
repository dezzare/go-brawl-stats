package configs

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
	var exist bool
	APIKey, exist = os.LookupEnv("APIKEY")
	if !exist {
		log.Printf("API Key não encontrada")
	}

}

const BaseURL = "https://api.brawlstars.com/v1"
const Port = "8080"

// const Tag = "%23V0CJ2J"
