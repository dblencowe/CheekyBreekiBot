package maps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var LoadedMaps []TarkovMap

type mapData struct {
	ID     int `json:"id"`
	Locale struct {
		En string `json:"en"`
	} `json:"locale"`
	Wiki         string   `json:"wiki"`
	Description  string   `json:"description"`
	Enemies      []string `json:"enemies"`
	RaidDuration struct {
		Day   int `json:"day"`
		Night int `json:"night"`
	} `json:"raidDuration"`
	Svg struct {
		File         string   `json:"file"`
		Floors       []string `json:"floors"`
		DefaultFloor string   `json:"defaultFloor"`
	} `json:"svg"`
	PlayerCount string `json:"playerCount"`
}

func LoadMaps() {
	jsonFile, err := os.Open("./data/maps.json")
	if err != nil {
		fmt.Println("Unable to open maps file", err)
		return
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var maps map[string]mapData
	json.Unmarshal([]byte(byteValue), &maps)
	for _, mapObject := range maps {
		LoadedMaps = append(LoadedMaps, TarkovMap{Id: mapObject.ID, Name: mapObject.Locale.En, Slug: strings.ToLower(mapObject.Locale.En), WikiUrl: mapObject.Wiki, RaidDuration: MapDuration(mapObject.RaidDuration), PlayerCount: mapObject.PlayerCount})
		fmt.Printf("Loaded Map %s from %+v\n", LoadedMaps[len(LoadedMaps)-1].Slug, mapObject)
	}
}

func GetMap(mapName string) *TarkovMap {
	for i := range LoadedMaps {
		if LoadedMaps[i].Slug == mapName {
			return &LoadedMaps[i]
		}
	}

	return nil
}
