package models

type Accessory struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type AccessoryList struct {
	List []Accessory `json:"name"`
	Id   int         `json:"id"`
}

type Brawler struct {
	Name       string `json:"name"`
	Gadgets    []Accessory
	Id         int `json:"id"`
	StarPowers []StarPower
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
