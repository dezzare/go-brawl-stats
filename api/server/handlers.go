package server

import (
	"net/http"

	apiclient "github.com/dezzare/go-brawl-stats/api/client"
)

func getPlayer(w http.ResponseWriter, r *http.Request) {
	playerTag := r.PathValue("playerTag")
	c := apiclient.New()
	c.GetPlayer(playerTag)

}
