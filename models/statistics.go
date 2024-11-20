package models

type Statistics struct {
	TotalMatches int                          `json:"totalMatches"`
	Brawler      map[string]StatisticsBrawler `json:"brawler"`
	Mode         map[string]int               `json:"mode"`
	Type         map[string]int               `json:"type"`
}

type StatisticsBrawler struct {
	MatchesPlayed  int               `json:"matchesPlayed"`
	TotalWins      int               `json:"totalWins"`
	MatchesAgainst map[string]Result `json:"matchesAgainst"`
}

type StatisticsMatch struct {
	Team1  Team   `json:"team1"`
	Team2  Team   `json:"team2"`
	Result string `json:"result"`
}

type Result struct {
	Win  int `json:"win"`
	Draw int `json:"draw"`
	Lost int `json:"lost"`
}

type MostPlayed struct {
	Brawler string `json:"brawler"`
	Played  int    `json:"played"`
}

type MostWin struct {
	Name string `json:"name"`
	Wins int    `json:"wins"`
}

type Team struct {
	P1 string `json:"p1"`
	P2 string `json:"p2"`
	P3 string `json:"p3"`
}

func (t *Team) AsStringSlice() []string {
	return []string{t.P1, t.P2, t.P3}
}
