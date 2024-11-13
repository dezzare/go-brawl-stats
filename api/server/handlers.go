package server

import (
	"fmt"
	"net/http"
	"strings"

	apiclient "github.com/dezzare/go-brawl-stats/api/client"
)

func getPlayer(w http.ResponseWriter, r *http.Request) {
	playerTag := "%23" + strings.TrimPrefix(r.PathValue("playerTag"), "#")
	c := apiclient.New()

	fmt.Printf("\nGetting player: %v", playerTag)
	err := c.GetPlayer(playerTag)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}

}

func getPlayersRankingsByCountry(w http.ResponseWriter, r *http.Request) {
	countryTag := r.PathValue("countryTag")
	if countryTag == "" {
		countryTag = "global"
	}

	c := apiclient.New()
	fmt.Printf("\nGetting Players Rank from: %v", countryTag)
	err := c.GetPlayersRankingsByCountry(countryTag)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
}
