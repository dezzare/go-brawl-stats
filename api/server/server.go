package server

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
)

func New() *http.Server {


	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/players/{playerTag}", getPlayer)
	mux.HandleFunc("GET /v1/players/{playerTag}/battlelog", getPlayerBattleLog)
	mux.HandleFunc("GET /v1/ranking/{countryTag}", getPlayersRankingsByCountry)
	mux.HandleFunc("GET /v1/brawlers", getBrawlers)
	mux.HandleFunc("GET /teste", teste)

	srv := &http.Server{
		Addr:    ":" + viper.GetString("Port"),
		Handler: mux,
	}

	return srv
}

func teste(w http.ResponseWriter, r *http.Request) {
	// c := client.New()
	// playerTag := "%23V0CJ2J"

	// _, err := c.GetPlayer(playerTag)
	// if err != nil {
	// 	fmt.Printf("\nError: %v", err)
	// }

	// _, err = c.GetPlayersRankingsByCountry("br")
	// if err != nil {
	// 	fmt.Printf("\nError: %v", err)
	// }

	// fmt.Println("Open file")
	// var file models.PlayerRankingList
	// err := getJsonFromFile("data-Top-Rank-Players.json", &file)
	// if err != nil {
	// 	fmt.Printf("\nError: %v", err)
	// }
	// fmt.Printf("\nFile Open: \n%v", file)

	// fmt.Println("\ntagList = getTags(file)")
	// tagList := getTopPlayersTags(file)
	// fmt.Printf("\nGet tags result: %v", tagList)

	// fmt.Println("\ngetAllPlayersBattleLog(tagList)")
	// getAllPlayersBattleLog(tagList)

	fmt.Println("Get stats:")
	extractStatics()

}
