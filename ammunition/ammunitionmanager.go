package ammunition

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dblencowe/CheekyBreekiBot/helper"
)

type TarkovAmmunition struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	ShortName       string  `json:"shortName"`
	Weight          float64 `json:"weight"`
	Caliber         string  `json:"caliber"`
	StackMaxSize    int     `json:"stackMaxSize"`
	Tracer          bool    `json:"tracer"`
	TracerColor     string  `json:"tracerColor"`
	AmmoType        string  `json:"ammoType"`
	ProjectileCount int     `json:"projectileCount"`
	Ballistics      struct {
		Damage              int     `json:"damage"`
		ArmorDamage         int     `json:"armorDamage"`
		FragmentationChance float64 `json:"fragmentationChance"`
		RicochetChance      float64 `json:"ricochetChance"`
		PenetrationChance   float64 `json:"penetrationChance"`
		PenetrationPower    int     `json:"penetrationPower"`
		Accuracy            int     `json:"accuracy"`
		Recoil              int     `json:"recoil"`
		InitialSpeed        int     `json:"initialSpeed"`
	} `json:"ballistics"`
}

var LoadedAmmunition []TarkovAmmunition
var LoadedCalibers []string

func LoadAmmunition(filePath string) {
	if len(filePath) == 0 {
		filePath = "./data/ammunition.json"
	}
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to open ammunition file", err)
		return
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var traders map[string]TarkovAmmunition
	json.Unmarshal([]byte(byteValue), &traders)
	for _, ammoObject := range traders {
		LoadedAmmunition = append(LoadedAmmunition, ammoObject)
		cleanCaliber := strings.Replace(ammoObject.Caliber, "Caliber", "", -1)
		_, contains := helper.Contains(LoadedCalibers, cleanCaliber)
		if !contains {
			LoadedCalibers = append(LoadedCalibers, cleanCaliber)
		}
	}
	fmt.Printf("Loaded %d ammunition records\n", len(LoadedAmmunition))
	fmt.Printf("Loaded %d calibers\nCalibers: %s\n", len(LoadedCalibers), strings.Join(LoadedCalibers, ", "))
}

func GetAmmoById(traderId int) *TarkovAmmunition {
	for i := range LoadedAmmunition {
		if LoadedAmmunition[i].ID == traderId {
			return &LoadedAmmunition[i]
		}
	}

	return nil
}

func GetAmmosByCaliber(caliber string) []TarkovAmmunition {
	var results []TarkovAmmunition
	for i := range LoadedAmmunition {
		if LoadedAmmunition[i].Caliber == "Caliber"+caliber {
			results = append(results, LoadedAmmunition[i])
		}
	}

	return results
}

type caliberComparison struct {
	Distance int
	Caliber  string
}

func GetCaliber(caliber string) string {
	var closestComparison caliberComparison
	closestComparison.Distance = 0
	for i := range LoadedCalibers {
		distance := helper.Levenshtein(LoadedCalibers[i], caliber)
		if distance <= closestComparison.Distance {
			closestComparison = caliberComparison{Distance: distance, Caliber: LoadedCalibers[i]}
		}
	}

	return closestComparison.Caliber
}
