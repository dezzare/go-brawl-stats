package server

import (
	"fmt"
	"net/http"

	"github.com/dezzare/go-brawl-stats/api/client"
	"github.com/dezzare/go-brawl-stats/configs"
)

func New() *http.Server {

	mux := http.NewServeMux()
	mux.HandleFunc("/", teste)
	mux.HandleFunc("GET /v1/players/{playerTag}", getPlayer)
	//mux.HandleFunc("/teste", teste)

	srv := &http.Server{
		Addr:    ":" + configs.Port,
		Handler: mux,
	}

	return srv
}

func teste(w http.ResponseWriter, r *http.Request) {
	c := client.New()
	tag := "%23V0CJ2J"
	err := c.GetPlayer(tag)
	if err != nil {
		fmt.Printf("Error hadler: %v", err)
	}

}
