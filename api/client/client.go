package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

	var jsonData models.Player
	fmt.Printf("\ndecoder:\n", data)
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return fmt.Errorf("Json Unmarshal error: %v", err)
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

func (c *Client) checkStatusCode(r *http.Response) {
	if r.StatusCode >= 200 && r.StatusCode < 300 {

	}
}
