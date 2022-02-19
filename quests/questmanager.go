package quests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var LoadedQuests []TarkovQuest

type TarkovQuest struct {
	ID      int `json:"id"`
	Require struct {
		Level  interface{} `json:"level"`
		Quests []int       `json:"quests"`
	} `json:"require"`
	Giver   int    `json:"giver"`
	Turnin  int    `json:"turnin"`
	Title   string `json:"title"`
	Locales struct {
		En string `json:"en"`
	} `json:"locales"`
	Nokappa    bool          `json:"nokappa"`
	Wiki       string        `json:"wiki"`
	Exp        int           `json:"exp"`
	Unlocks    []interface{} `json:"unlocks"`
	Reputation []struct {
		Trader int     `json:"trader"`
		Rep    float64 `json:"rep"`
	} `json:"reputation"`
	Objectives []struct {
		Type     string `json:"type"`
		Target   string `json:"target"`
		Number   int    `json:"number"`
		Location int    `json:"location"`
		ID       int    `json:"id"`
	} `json:"objectives"`
	GameID string `json:"gameId"`
}

func LoadQuests() {
	jsonFile, err := os.Open("./data/quests.json")
	if err != nil {
		fmt.Println("Unable to open quests file", err)
		return
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var quests []TarkovQuest
	json.Unmarshal([]byte(byteValue), &quests)
	for _, questObject := range quests {
		LoadedQuests = append(LoadedQuests, questObject)
	}
	fmt.Printf("Loaded %d quests\n", len(LoadedQuests))
}

func GetQuest(questName string) *TarkovQuest {
	for i := range LoadedQuests {
		if strings.ToLower(LoadedQuests[i].Title) == questName {
			return &LoadedQuests[i]
		}
	}

	return nil
}
