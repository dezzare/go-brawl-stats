package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dezzare/go-brawl-stats/config"
)

func init() {
	log.Println("Loading .env File")
	config.LoadEnvFile()
}
func main() {

	C = newClient(config.APIKey, config.BaseURL)

	mux := http.NewServeMux()
	mux.HandleFunc("/", teste)
	//mux.HandleFunc("/teste", teste)

	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: mux,
	}

	log.Printf("Server is running on port: %v\n", config.Port)
	log.Fatal(srv.ListenAndServe())
}

func teste(w http.ResponseWriter, r *http.Request) {
	//c := newClient(config.APIKey, config.BaseURL)
	tag := config.Tag
	err := C.GetPlayer(tag)
	if err != nil {
		fmt.Printf("Error hadler: %v", err)
	}

}
