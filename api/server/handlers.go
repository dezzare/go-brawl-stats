package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	apiclient "github.com/dezzare/go-brawl-stats/api/client"
	"github.com/dezzare/go-brawl-stats/models"
)

func getPlayer(w http.ResponseWriter, r *http.Request) {
	playerTag := parseTag(r.PathValue("playerTag"))
	c := apiclient.New()

	fmt.Printf("\nGetting player: %v", playerTag)
	player, err := c.GetPlayer(playerTag)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	var model models.Player
	saveToJsonFile(player, model, "data-Player.json")

}

func getBrawlers(w http.ResponseWriter, r *http.Request) {
	c := apiclient.New()

	fmt.Println("\nGetting List of Brawlers")
	brawlers, err := c.GetBrawlers()
	if err != nil {
		fmt.Printf("%v", err)
	}
	var model models.BrawlerList
	saveToJsonFile(brawlers, model, "data-Brawlers-List.json")
}

func getAllPlayers(tagList []string) {
	c := apiclient.New()
	var list models.PlayerList
	for k, v := range tagList {
		fmt.Printf("\n%v - Get player with id: %v", k, v)
		p, _ := c.GetPlayer(v)
		var aux models.Player
		if err := json.Unmarshal(p, &aux); err != nil {
			fmt.Printf("Json Unmarshal error: %v", err)
		}
		fmt.Printf("\nPlayer appended: \n%v", aux)
		fmt.Println("--------")
		list.Player = append(list.Player, aux)
	}
	var model models.PlayerList
	data, err := json.Marshal(list)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	saveToJsonFile(data, model, "data-Players-List.json")
}

func getAllPlayersBattleLog(tagList []string) {
	c := apiclient.New()
	var battles models.AllBattles
	for k, v := range tagList {
		fmt.Printf("\n%v - Get player with id: %v", k, v)
		b, _ := c.GetPlayerBattleLog(v)
		var aux models.BattleList
		if err := json.Unmarshal(b, &aux); err != nil {
			fmt.Printf("Json Unmarshal error: %v", err)
		}
		fmt.Printf("\nBattle appended: \n%v", aux)
		fmt.Println("--------")
		battles.List = battles.Add(aux)
		// battles.Add(aux)
	}
	var model models.AllBattles
	data, err := json.Marshal(battles)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	saveToJsonFile(data, model, "data-AllPlayers-BattleLog.json")
}

func getPlayersRankingsByCountry(w http.ResponseWriter, r *http.Request) {
	countryTag := r.PathValue("countryTag")
	if countryTag == "" {
		countryTag = "global"
	}

	c := apiclient.New()
	fmt.Printf("\nGetting Players Rank from: %v", countryTag)
	players, err := c.GetPlayersRankingsByCountry(countryTag)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}

	var model models.PlayerRankingList
	saveToJsonFile(players, model, "data-Top-Rank-Players.json")
}

func getPlayerBattleLog(w http.ResponseWriter, r *http.Request) {
	playerTag := parseTag(r.PathValue("playerTag"))
	c := apiclient.New()

	fmt.Printf("\nGetting Player battlelog: %v", playerTag)
	player, err := c.GetPlayerBattleLog(playerTag)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	var model models.BattleList
	saveToJsonFile(player, model, "data-Player-BattleLog.json")

}
