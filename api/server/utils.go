package server

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/dezzare/go-brawl-stats/models"
)

func getJsonFromFile[T interface{}](filename string, model *T) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Open file error: %v", err)
	}
	if err := json.Unmarshal(file, &model); err != nil {
		return fmt.Errorf("Json Unmarshal error: %v", err)
	}
	return nil
}

func getTopPlayersTags(list models.PlayerRankingList) []string {
	values := list.PlayerRanking
	var tags []string
	for i := 0; i < len(values); i++ {
		tags = append(tags, parseTag(values[i].Tag))
	}
	fmt.Printf("\nTags: %v", tags)
	return tags
}

func getBrawlersResult() map[string]models.Result {
	bl := getBrawlerList()
	bm := make(map[string]models.Result)
	r := models.Result{
		Win:  0,
		Draw: 0,
		Lost: 0,
	}
	for i := 0; i < len(bl.Brawler); i++ {
		bm[bl.Brawler[i].Name] = r
	}
	return bm
}

func getBrawlerList() models.BrawlerList {
	var bl models.BrawlerList
	if err := getJsonFromFile("data-Brawlers-List.json", &bl); err != nil {
		fmt.Print(err)
	}
	return bl
}

func getBrawlersStatsInterface() map[string]models.StatisticsBrawler {
	bl := getBrawlerList()
	bm := getBrawlersResult()
	m := make(map[string]models.StatisticsBrawler)
	for i := 0; i < len(bl.Brawler); i++ {
		b := models.StatisticsBrawler{
			MatchesPlayed:  0,
			TotalWins:      0,
			MatchesAgainst: bm,
		}

		m[bl.Brawler[i].Name] = b
	}

	return m
}

func extractStatics() {

	stats := extractBattleStats()

	fmt.Println("\n----------------")
	fmt.Println("Statistics")
	fmt.Println("----------------")

	fmt.Printf("\nTotal Matches: %v", stats.TotalMatches)

	for k, v := range stats.Mode {
		fmt.Printf("\nMode: %v, was played %v times", k, v)
	}

	for k, v := range stats.Type {
		fmt.Printf("\nType: %v, was played %v times", k, v)
	}

	mp := extractMostPlayedBrawler(stats)
	fmt.Printf("\nMost Played Brawler: %v, played for %v times", mp.Brawler, mp.Played)

	mw := extractMostWinBrawler(stats)
	fmt.Printf(`
	Pos.	Brawler		Victories
	1 - 	%v			%v
	2 - 	%v			%v
	3 - 	%v			%v
	4 - 	%v			%v
	5 - 	%v			%v

		`, mw[0].Name, mw[0].Wins,
		mw[1].Name, mw[1].Wins,
		mw[2].Name, mw[2].Wins,
		mw[3].Name, mw[3].Wins,
		mw[4].Name, mw[4].Wins,
	)
}

// type Statistics struct {
// 	TotalMatches int                          `json:"totalMatches"`
// 	Brawler      map[string]StatisticsBrawler `json:"brawler"`
// 	Mode         map[string]int               `json:"mode"`
// 	Type         map[string]int               `json:"type"`
// }

// type StatisticsBrawler struct {
// 	MatchesPlayed int               `json:"matchesPlayed"`
// 	TotalWins     int               `json:"totalWins"`
// 	MachesAgainst map[string]Result `json:"machesAgainst"`

func extractBattleStats() models.Statistics {
	var data models.AllBattles
	if err := getJsonFromFile("data-AllPlayers-BattleLog.json", &data); err != nil {
		fmt.Print(err)
	}

	s := models.Statistics{
		TotalMatches: 0,
		Brawler:      getBrawlersStatsInterface(),
		Mode:         make(map[string]int),
		Type:         make(map[string]int),
	}
	// for i := 0; i < len(data.List); i++ {
	// 	for j := 0; j < len(data.List[i].Battle); j++ {

	// 		s.TotalMatches++
	// 		s.Mode[data.List[i].Battle[j].BattleResult.Mode]++
	// 		s.Type[data.List[i].Battle[j].BattleResult.Type]++

	// 		bn, _ := extractBrawlerNames(data.List[i].Battle[j].BattleResult)
	// 		bn.Result = data.List[i].Battle[j].BattleResult.Result
	// 		s = extractBrawlerStats(s, bn)
	// 	}
	for _, v := range data.List {
		for _, vv := range v.Battle {
			if checkBattleMap(vv.Event.Mode) {

				s.TotalMatches++
				s.Mode[vv.BattleResult.Mode]++
				s.Type[vv.BattleResult.Type]++

				bn, _ := extractBrawlerNames(vv.BattleResult)
				bn.Result = vv.BattleResult.Result

				s = extractBrawlerStats(s, bn)

			}

		}
	}
	return s
}

func extractBrawlerNames(battle models.BattleResult) (sbt models.StatisticsMatch, err error) {
	count := 0
	// sbt.Team2 = append(sbt.Team2, battle.Teams[1][2].Brawler.Name)

	for _, v := range battle.Teams {
		for i := 0; i < len(v); i++ {
			switch {
			case count == 0:
				sbt.Team1.P1 = v[i].Brawler.Name

			case count == 1:
				sbt.Team1.P2 = v[i].Brawler.Name
			case count == 2:
				sbt.Team1.P3 = v[i].Brawler.Name

			case count == 3:
				sbt.Team2.P1 = v[i].Brawler.Name
			case count == 4:
				sbt.Team2.P2 = v[i].Brawler.Name
			case count == 5:
				sbt.Team2.P3 = v[i].Brawler.Name
			}
			count++
		}

	}
	// fmt.Printf(`
	// 	Type: %v
	// 	Count: %v
	// 	SBT.Team1: %v
	// 	SBT.Team2: %v
	// 	`, battle.Type, count, sbt.Team1, sbt.Team2)

	return
}

func checkBattleMap(s string) bool {
	acceptedMaps := []string{"heist", "bounty", "gemGrab", "brawlBall", "knockout", "hotZone"}
	for _, v := range acceptedMaps {
		if v == s {
			return true
		}
	}
	return false
}

func extractBrawlerStats(s models.Statistics, bn models.StatisticsMatch) models.Statistics {
	switch bn.Result {
	case "victory":
		s = addWins(s, bn.Team1.AsStringSlice(), bn.Team2.AsStringSlice())
		s = addLost(s, bn.Team1.AsStringSlice(), bn.Team2.AsStringSlice())
	case "defeat":
		s = addWins(s, bn.Team2.AsStringSlice(), bn.Team1.AsStringSlice())
		s = addLost(s, bn.Team2.AsStringSlice(), bn.Team1.AsStringSlice())
	default:
		s = addDraw(s, bn.Team1.AsStringSlice(), bn.Team2.AsStringSlice())
		s = addDraw(s, bn.Team2.AsStringSlice(), bn.Team1.AsStringSlice())
	}
	return s
}

func addWins(s models.Statistics, wTeam, lTeam []string) models.Statistics {
	for i := 0; i < len(wTeam); i++ {
		for j := 0; j < len(lTeam); j++ {
			aux := models.Result{
				Win:  s.Brawler[wTeam[i]].MatchesAgainst[lTeam[j]].Win + 1,
				Draw: s.Brawler[wTeam[i]].MatchesAgainst[lTeam[j]].Draw,
				Lost: s.Brawler[wTeam[i]].MatchesAgainst[lTeam[j]].Lost,
			}
			s.Brawler[wTeam[i]].MatchesAgainst[lTeam[j]] = aux
		}
		aux := models.StatisticsBrawler{
			TotalWins:      s.Brawler[wTeam[i]].TotalWins + 1,
			MatchesPlayed:  s.Brawler[wTeam[i]].MatchesPlayed + 1,
			MatchesAgainst: s.Brawler[wTeam[i]].MatchesAgainst,
		}
		s.Brawler[wTeam[i]] = aux
	}
	return s
}

func addLost(s models.Statistics, wTeam, lTeam []string) models.Statistics {
	for i := 0; i < len(lTeam); i++ {
		for j := 0; j < len(wTeam); j++ {
			aux := models.Result{
				Win:  s.Brawler[lTeam[i]].MatchesAgainst[wTeam[j]].Win,
				Draw: s.Brawler[lTeam[i]].MatchesAgainst[wTeam[j]].Draw,
				Lost: s.Brawler[lTeam[i]].MatchesAgainst[wTeam[j]].Lost + 1,
			}
			s.Brawler[wTeam[i]].MatchesAgainst[lTeam[j]] = aux
		}
		aux := models.StatisticsBrawler{
			TotalWins:      s.Brawler[lTeam[i]].TotalWins,
			MatchesPlayed:  s.Brawler[lTeam[i]].MatchesPlayed + 1,
			MatchesAgainst: s.Brawler[lTeam[i]].MatchesAgainst,
		}
		s.Brawler[lTeam[i]] = aux
	}
	return s
}

func addDraw(s models.Statistics, team1, team2 []string) models.Statistics {
	for i := 0; i < len(team1); i++ {
		for j := 0; j < len(team2); j++ {
			aux1 := models.Result{
				Win:  s.Brawler[team1[i]].MatchesAgainst[team2[j]].Win,
				Draw: s.Brawler[team1[i]].MatchesAgainst[team2[j]].Draw + 1,
				Lost: s.Brawler[team1[i]].MatchesAgainst[team2[j]].Lost,
			}
			s.Brawler[team1[i]].MatchesAgainst[team2[j]] = aux1

		}
		aux2 := models.StatisticsBrawler{
			TotalWins:      s.Brawler[team1[i]].TotalWins,
			MatchesPlayed:  s.Brawler[team1[i]].MatchesPlayed + 1,
			MatchesAgainst: s.Brawler[team1[i]].MatchesAgainst,
		}
		s.Brawler[team1[i]] = aux2
	}
	return s
}

func extractMostPlayedBrawler(s models.Statistics) models.MostPlayed {

	var mp []models.MostPlayed
	for k, v := range s.Brawler {
		mp = append(mp, models.MostPlayed{Brawler: k, Played: v.MatchesPlayed})
	}
	sort.Sort(ByMatchesPlayed(mp))

	return mp[0]
}

func extractMostWinBrawler(s models.Statistics) []models.MostWin {
	var mw []models.MostWin
	for k, v := range s.Brawler {
		mw = append(mw, models.MostWin{Name: k, Wins: v.TotalWins})
	}
	sort.Sort(ByMostWins(mw))

	return mw
}

func parseTag(tag string) string {
	return "%23" + strings.TrimPrefix(strings.TrimPrefix(tag, "#"), "%23")
}

func saveToJsonFile[T interface{}](data []byte, model T, filename string) error {

	if err := json.Unmarshal(data, &model); err != nil {
		return fmt.Errorf("Json Unmarshal error: %v", err)
	}

	file, err := json.MarshalIndent(&model, "", " ")
	if err != nil {
		return fmt.Errorf("Json Marshal error: %v", err)
	}
	err = os.WriteFile(filename, file, 0644)
	if err != nil {
		return fmt.Errorf("Write error: %v", err)
	}
	return nil
}
