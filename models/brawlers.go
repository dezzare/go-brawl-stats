package models

type Accessory struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

// TODO remove?
type AccessoryList struct {
	List []Accessory `json:"name"`
	Id   int         `json:"id"`
}

type Brawler struct {
	Name       string      `json:"name"`
	Gadgets    []Accessory `json:"gadgets"`
	StarPowers []StarPower `json:"starPowers"`
	Id         int         `json:"id"`
}

type BrawlerList struct {
	Brawler []Brawler `json:"items"`
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

type StarPower struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}
