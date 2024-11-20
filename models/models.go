package models

type AllBattles struct {
	List []BattleList `json:"list"`
}

func (b *AllBattles) Add(item BattleList) []BattleList {
	b.List = append(b.List, item)
	return b.List
}

type Battle struct {
	BattleTime   string       `json:"battleTime"`
	Event        Event        `json:"event"`
	BattleResult BattleResult `json:"battle"`
}

type BattleList struct {
	Battle []Battle `json:"items"`
}

type BattleLogPlayer []struct {
	Tag     string           `json:"tag"`
	Name    string           `json:"name"`
	Brawler BattleLogBrawler `json:"brawler"`
}

type BattleLogStarPlayer struct {
	Tag     string           `json:"tag"`
	Name    string           `json:"name"`
	Brawler BattleLogBrawler `json:"brawler"`
}

type BattleLogBrawler struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Power    int    `json:"power"`
	Trophies int    `json:"trophies"`
}

type BattleResult struct {
	Mode         string              `json:"mode"`
	Type         string              `json:"type"`
	Result       string              `json:"result"`
	Duration     int                 `json:"duration"`
	TrophyChange int                 `json:"trophyChange"`
	StarPlayer   BattleLogStarPlayer `json:"starPlayer"`
	Teams        []BattleLogPlayer   `json:"teams"`
}

type Event struct {
	Id   int    `json:"id"`
	Mode string `json:"mode"`
	Map  string `json:"map"`
}

type Player struct {
	Club                 PlayerClub    `json:"club"`
	IsPro                bool          `json:"isQualifiedFromChampionshipChallenge"`
	Win3x3               int           `json:"3vs3Victories"`
	Icon                 PlayerIcon    `json:"icon"`
	Tag                  string        `json:"tag"`
	Name                 string        `json:"name"`
	Trophies             int           `json:"trophies"`
	ExpLevel             int           `json:"expLevel"`
	ExpPoints            int           `json:"expPoints"`
	HighestTrophies      int           `json:"highestTrophies"`
	SoloVictories        int           `json:"soloVictories"`
	DuoVictories         int           `json:"duoVictories"`
	BestRoboRumbleTime   int           `json:"bestRoboRumbleTime"`
	BestTimeAsBigBrawler int           `json:"bestTimeAsBigBrawler"`
	Brawlers             []BrawlerStat `json:"brawlers"`
	NameColor            string        `json:"nameColor"`
}

type PlayerList struct {
	Player []Player `json:"player"`
}

type PlayerClub struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

type PlayerIcon struct {
	Id int `json:"id"`
}

type PlayerRankingList struct {
	PlayerRanking []PlayerRanking `json:"items"`
}
type PlayerRanking struct {
	Club      PlayerClub `json:"club"`
	Icon      PlayerIcon `json:"icon"`
	Trophies  int        `json:"trophies"`
	Tag       string     `json:"tag"`
	Name      string     `json:"name"`
	Rank      int        `json:"rank"`
	NameColor string     `json:"nameColor"`
}

// type Statistics struct {
// 	TotalMatches int
// 	Mode         map[string]int
// 	Type         map[string]int
// 	Brawler      map[string]StatisticsBrawler struct {
// 						MatchesPlayed int
// 						TotalWins     int
// 						MachesAgainst map[string]Result{
// 												Win  int
// 												Draw int
// 												Lost int
// 						}
// }
// }
//
// type AllBattles struct {
// List []BattleList
// 		Battle []Battle
// 			BattleTime   string
// 			Event        Event
// 			BattleResult BattleResult
// 				Mode         string
// 				Type         string
// 				Result       string
// 				Duration     int
// 				TrophyChange int
// 				StarPlayer   BattleLogStarPlayer
// 				Teams        []BattleLogPlayer
// 					Tag     string
// 					Name    string
// 					Brawler BattleLogBrawler
// 						Id       int
// 						Name     string
// 						Power    int
// 						Trophies int
//}
