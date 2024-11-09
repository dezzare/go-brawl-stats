package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dezzare/go-brawl-stats/config"
)

const userTag = "%23V0CJ2J"

func init() {
	log.Println("Loading .env File")
	config.LoadEnvFile()
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/teste", teste)

	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: mux,
	}

	log.Printf("Server is running on port: %v\n", config.Port)
	log.Fatal(srv.ListenAndServe())
}

func teste(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("\nAPI k: %v", config.APIKey)
	fmt.Printf("\nTest: %s\n", config.BaseURL)
}

func handler(w http.ResponseWriter, r *http.Request) {

	req, _ := http.NewRequest("GET", config.BaseURL+"/players/"+userTag, nil)
	req.Header.Set("Authorization", "Bearer "+config.APIKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("\nHttp request error: %v", err)
	}
	defer res.Body.Close()

	var jsonData map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&jsonData)
	if err != nil {
		log.Printf("Json Unmarshal error: %v", err)
	}

	log.Printf("\njson data: \n%v", jsonData)

}
