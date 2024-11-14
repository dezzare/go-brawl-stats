package server

import (
	"fmt"
	"net/http"

	"github.com/dezzare/go-brawl-stats/api/client"
	"github.com/dezzare/go-brawl-stats/configs"
	"github.com/dezzare/go-brawl-stats/models"
)

func New() *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/players/{playerTag}", getPlayer)
	mux.HandleFunc("GET /v1/players/{playerTag}/battlelog", getPlayerBattleLog)
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
	_, err := c.GetPlayer(playerTag)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}

	fmt.Println("\nGetting Players Rank")
	_, err = c.GetPlayersRankingsByCountry("br")
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}

	fmt.Println("Open file")
	var file models.PlayerRankingList
	err = getJsonFromFile("test-PlayerRankingList.json", &file)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	fmt.Printf("\nFile Open: \n%v", file)

	fmt.Println("\ntagList = getTags(file)")
	tagList := getTags(file)
	fmt.Printf("\nGet tags result: %v", tagList)

	fmt.Println("\ngetAllPlayersBattleLog(tagList)")
	getAllPlayersBattleLog(tagList)

}
