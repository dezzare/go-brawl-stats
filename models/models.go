package models

type Accessory struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type AccessoryList struct {
	List []Accessory `json:"name"`
	Id   int         `json:"id"`
}

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

type Brawler struct {
	Gadgets    []Accessory `json:"gadgets"`
	Name       string      `json:"name"`
	Id         int         `json:"id"`
	StarPowers []StarPower `json:"starPowers"`
}

type BrawlerStat struct {
	Name            string      `json:"name"`
	Gadgets         []Accessory `json:"gadgets"`
	StarPowers      []StarPower `json:"starPowers"`
	Id              int         `json:"id"`
	Rank            int         `json:"rank"`
	Trophies        int         `json:"trophies"`
	HighestTrophies int         `json:"highestTrophies"`
	Power           int         `json:"power"`
	Gears           []GearStat  `json:"gears"`
}

type Event struct {
	Id   int    `json:"id"`
	Mode string `json:"mode"`
	Map  string `json:"map"`
}

type GearInfo struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Level int    `json:"level"`
}

type GearStat struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Level int    `json:"level"`
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

type StarPower struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
