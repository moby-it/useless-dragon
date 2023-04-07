package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var config Config

type Buff struct {
	Duration int `json:"duration"`
	Attack   int `json:"attack"`
	Defence  int `json:"defence"`
}
type Config struct {
	Buffs            map[string]Buff `json:"buffs"`
	PowerAttackBonus int             `json:"power attack bonus"`
	Guard_Bonus      int             `json:"guard defence bonus"`
}

func Get() Config {
	return config
}
func Parse() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(fmt.Sprintf("%v/assets/config.json", dir))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}
