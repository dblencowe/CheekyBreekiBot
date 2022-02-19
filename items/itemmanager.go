package items

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TarkovItem struct {
	Id        string
	Name      string
	Slug      string
	ShortName string
}

var LoadedItems []TarkovItem

type itemData struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

func LoadItems() {
	jsonFile, err := os.Open("./data/items.json")
	if err != nil {
		fmt.Println("Unable to open items file", err)
		return
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var items map[string]itemData
	json.Unmarshal([]byte(byteValue), &items)
	for _, itemObject := range items {
		LoadedItems = append(LoadedItems, TarkovItem{
			Id:        itemObject.Id,
			Name:      itemObject.Name,
			ShortName: itemObject.ShortName,
		})
	}
	fmt.Printf("Loaded %d items into memory\n", len(LoadedItems))
}

func GetItemById(itemId string) *TarkovItem {
	for i := range LoadedItems {
		if LoadedItems[i].Id == itemId {
			return &LoadedItems[i]
		}
	}

	return nil
}
