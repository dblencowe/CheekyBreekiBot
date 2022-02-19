package maps

import "fmt"

type MapDuration struct {
	Day   int
	Night int
}

type TarkovMap struct {
	Id           int `json:"id"`
	Slug         string
	Name         string `json:"locale.en"`
	WikiUrl      string `json:"wiki"`
	RaidDuration MapDuration
	PlayerCount  string
}

func (m TarkovMap) Summary() string {
	return fmt.Sprintf("%s has %s players and lasts %d minutes during the day and %d minutes at night", m.Name, m.PlayerCount, m.RaidDuration.Day, m.RaidDuration.Night)
}
