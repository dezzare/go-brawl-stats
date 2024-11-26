package main

import (
	"log"

	"github.com/dezzare/go-brawl-stats/api/server"
	"github.com/dezzare/go-brawl-stats/configs"
	"github.com/spf13/viper"
)

func main() {

	//cmd.Execute()
	configs.Load()
	srv := server.New()

	log.Printf("Server is running on port: %s\n", viper.GetString("Port"))
	log.Fatal(srv.ListenAndServe())
}
