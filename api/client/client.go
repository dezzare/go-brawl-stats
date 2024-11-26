package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

type Client struct {
	HTTP    *http.Client
	BaseURL string
	APIKey  string
}

func New() *Client {

	return &Client{
		HTTP:    &http.Client{},
		BaseURL: viper.GetString("BaseURL"),
		APIKey:  viper.GetString("APIKey"),
	}
}

func (c *Client) GetPlayer(playerTag string) (data []byte, err error) {
	data, err = c.doRequest("GET", c.BaseURL+"/players/"+playerTag)
	if err != nil {
		return data, fmt.Errorf("Error getting player: %v", err)
	}
	return
}

func (c *Client) GetBrawlers() (data []byte, err error) {
	data, err = c.doRequest("GET", c.BaseURL+"/brawlers")
	if err != nil {
		return data, fmt.Errorf("Error getting brawlers: %v", err)
	}
	return
}

func (c *Client) GetPlayerBattleLog(playerTag string) (data []byte, err error) {
	data, err = c.doRequest("GET", c.BaseURL+"/players/"+playerTag+"/battlelog")
	if err != nil {
		return data, fmt.Errorf("Error getting player battlelog: %v", err)
	}
	return
}

func (c *Client) GetPlayersRankingsByCountry(countryCode string) (data []byte, err error) {
	data, err = c.doRequest("GET", c.BaseURL+"/rankings/"+countryCode+"/players")
	if err != nil {
		return data, fmt.Errorf("Error getting Players Ranings by Country: %v", err)
	}
	return
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
	fmt.Printf("\nStatus Code: %v\nHTTP Status: %v", res.StatusCode, http.StatusText(res.StatusCode))
	return
}
