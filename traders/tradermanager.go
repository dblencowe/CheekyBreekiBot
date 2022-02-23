package traders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var LoadedTraders []TarkovTrader

type TarkovTrader struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Locale struct {
		En string `json:"en"`
	} `json:"locale"`
	Wiki          string   `json:"wiki"`
	Description   string   `json:"description"`
	Currencies    []string `json:"currencies"`
	SalesCurrency string   `json:"salesCurrency"`
	Loyalty       []struct {
		Level              int `json:"level"`
		RequiredLevel      int `json:"requiredLevel"`
		RequiredReputation int `json:"requiredReputation"`
		RequiredSales      int `json:"requiredSales"`
	} `json:"loyalty"`
}

func LoadTraders() {
	jsonFile, err := os.Open("./data/traders.json")
	if err != nil {
		fmt.Println("Unable to open traders file", err)
		return
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var traders map[string]TarkovTrader
	json.Unmarshal([]byte(byteValue), &traders)
	for _, traderObject := range traders {
		LoadedTraders = append(LoadedTraders, traderObject)
	}
}

func GetTraderById(traderId int) *TarkovTrader {
	for i := range LoadedTraders {
		if LoadedTraders[i].ID == traderId {
			return &LoadedTraders[i]
		}
	}

	return nil
}
