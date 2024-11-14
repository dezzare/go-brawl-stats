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
	saveToJsonFile(player, model, "test-Player.json")

}

func getAllPlayers(tagList []string) {
	c := apiclient.New()
	var players models.PlayerList
	for k, v := range tagList {
		fmt.Printf("\n%v - Get player with id: %v", k, v)
		p, _ := c.GetPlayer(v)
		var model models.Player
		if err := json.Unmarshal(p, &model); err != nil {
			fmt.Printf("Json Unmarshal error: %v", err)
		}
		fmt.Printf("\nPlayer appended: \n%v", model)
		fmt.Println("--------")
		players.Player = append(players.Player, model)
	}
	var model models.PlayerList
	data, err := json.Marshal(players)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	saveToJsonFile(data, model, "test-Player-List.json")
}

func getAllPlayersBattleLog(tagList []string) {
	c := apiclient.New()
	var battles models.AllBattles
	for k, v := range tagList {
		fmt.Printf("\n%v - Get player with id: %v", k, v)
		b, _ := c.GetPlayerBattleLog(v)
		var model models.BattleList
		if err := json.Unmarshal(b, &model); err != nil {
			fmt.Printf("Json Unmarshal error: %v", err)
		}
		fmt.Printf("\nBattle appended: \n%v", model)
		fmt.Println("--------")
		battles.Add(model)

	}
	var model models.AllBattles
	data, err := json.Marshal(battles)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	saveToJsonFile(data, model, "test-AllPlayers-BattleLog.json")
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
	saveToJsonFile(players, model, "test-Top-Rank-Players.json")
}

func getPlayerBattleLog(w http.ResponseWriter, r *http.Request) {
	playerTag := parseTag(r.PathValue("playerTag"))
	c := apiclient.New()

	fmt.Printf("\nGetting Player battlelog: %v", playerTag)
	player, err := c.GetPlayer(playerTag)
	if err != nil {
		fmt.Printf("\nError: %v", err)
	}
	var model models.BattleList
	saveToJsonFile(player, model, "test-Player-BattleLog.json")

}
