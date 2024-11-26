package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	APIKey  string
	BaseURL string
	Port    string
}

// var APIKey string

// func LoadEnvFile() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	var exist bool
// 	APIKey, exist = os.LookupEnv("APIKEY")
// 	if !exist {
// 		log.Printf("API Key não encontrada")
// 	}

// }

// const BaseURL = "https://api.brawlstars.com/v1"
// const Port = "8080"

// const Tag = "%23V0CJ2J"

func Load() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viper.SetDefault("BaseURL", "https://api.brawlstars.com/v1")
	viper.SetDefault("Port", "8080")

	return
}
