package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	HTTP    *http.Client
	BaseURL string
	APIKey  string
}

var C *Client

func newClient(token string, baseURL string) *Client {
	return &Client{
		HTTP:    &http.Client{},
		BaseURL: baseURL,
		APIKey:  token,
	}
}

func (c *Client) GetPlayer(userTag string) error {
	err := c.doRequest("GET", c.BaseURL+"/players/"+userTag)
	if err != nil {
		return fmt.Errorf("Error getting player: %v", err)
	}
	return nil
}

func (c *Client) doRequest(method string, path string) error {
	req, _ := http.NewRequest(method, path, nil)
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	res, err := c.HTTP.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request error: %v", err)
	}
	defer res.Body.Close()

	var jsonData map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&jsonData)
	if err != nil {
		log.Printf("Json Unmarshal error: %v", err)
	}

	fmt.Printf("\njson data: \n%v", jsonData)

	return nil
}
