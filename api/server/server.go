package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/dezzare/go-brawl-stats/api/client"
	"github.com/dezzare/go-brawl-stats/configs"
	"github.com/dezzare/go-brawl-stats/models"
)

func New() *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/players/{playerTag}", getPlayer)
	mux.HandleFunc("GET /v1/ranking/{countryTag}", getPlayersRankingsByCountry)
	mux.HandleFunc("GET /teste", teste)

	srv := &http.Server{
		Addr:    ":" + configs.Port,
		Handler: mux,
	}

	return srv
}

func teste(w http.ResponseWriter, r *http.Request) {
	c := client.New()
	playerTag := "%23V0CJ2J"

	fmt.Println("\nGet player")
	err := c.GetPlayer(playerTag)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}

	fmt.Println("\nGetting Players Rank")
	err = c.GetPlayersRankingsByCountry("br")
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}

	fmt.Println("Open file")
	var file models.PlayerRankingList
	err = getJsonFromFile("test-PlayerRankingList.json", &file)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}

	fmt.Printf("\nFile:\n %v", file)

}

func getJsonFromFile[T interface{}](filename string, model *T) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Open file error: %v", err)
	}
	if err := json.Unmarshal(file, model); err != nil {
		return fmt.Errorf("Json Unmarshal error: %v", err)
	}
	return nil
}
