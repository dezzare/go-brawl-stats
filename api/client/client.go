package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dezzare/go-brawl-stats/configs"
	"github.com/dezzare/go-brawl-stats/models"
)

type Client struct {
	HTTP    *http.Client
	BaseURL string
	APIKey  string
}

func New() *Client {
	return &Client{
		HTTP:    &http.Client{},
		BaseURL: configs.BaseURL,
		APIKey:  configs.APIKey,
	}
}

func (c *Client) GetPlayer(playerTag string) error {
	data, err := c.doRequest("GET", c.BaseURL+"/players/"+playerTag)
	if err != nil {
		return fmt.Errorf("Error getting player: %v", err)
	}

	var player models.Player
	err = saveToJsonFile(data, &player, "test-Player.json")
	if err != nil {
		return fmt.Errorf("Erro Save to file: %v", err)
	}
	// var jsonData models.Player
	// if err := json.Unmarshal(data, &jsonData); err != nil {
	// 	return fmt.Errorf("Json Unmarshal error: %v", err)
	// }

	return nil
}

func (c *Client) GetPlayersRankingsByCountry(countryCode string) error {

	data, err := c.doRequest("GET", c.BaseURL+"/rankings/"+countryCode+"/players")
	if err != nil {
		return fmt.Errorf("Error getting Players Ranings by Country: %v", err)
	}

	var jsonData models.PlayerRankingList
	err = saveToJsonFile(data, &jsonData, "test-PlayerRankingList.json")
	if err != nil {
		return fmt.Errorf("Erro Save to file: %v", err)
	}

	return nil
}

func (c *Client) doRequest(method string, path string) (data []byte, err error) {
	req, _ := http.NewRequest(method, path, nil)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Read IO error: %v", err)
	}
	return
}

func saveToJsonFile[T interface{}](data []byte, model *T, filename string) error {

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
