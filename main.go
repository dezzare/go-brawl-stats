package main

import (
	"log"

	"github.com/dezzare/go-brawl-stats/api/server"
	"github.com/dezzare/go-brawl-stats/configs"
)

func init() {
	log.Println("Loading .env File")
	configs.LoadEnvFile()
}

func main() {

	srv := server.New()

	log.Printf("Server is running on port: %v\n", configs.Port)
	log.Fatal(srv.ListenAndServe())
}
