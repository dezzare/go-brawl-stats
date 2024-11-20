package server

import "github.com/dezzare/go-brawl-stats/models"

type ByMatchesPlayed []models.MostPlayed

func (w ByMatchesPlayed) Len() int {
	return len(w)
}

func (w ByMatchesPlayed) Less(i, j int) bool {
	return w[i].Played < w[j].Played
}

func (w ByMatchesPlayed) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

type ByMostWins []models.MostWin

func (w ByMostWins) Len() int {
	return len(w)
}

func (w ByMostWins) Less(i, j int) bool {
	return w[i].Wins < w[j].Wins
}

func (w ByMostWins) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
