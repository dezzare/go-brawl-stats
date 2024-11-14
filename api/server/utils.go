package server

import (
	"encoding/json"
	"fmt"
	"os"
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

func getTags(model models.PlayerRankingList) []string {
	values := model.PlayerRanking
	var tags []string
	for i := 0; i < len(values); i++ {
		tags = append(tags, parseTag(values[i].Tag))
	}
	fmt.Printf("\nTags: %v", tags)
	return tags
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
