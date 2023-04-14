package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"useless_dragon/combat"
	"useless_dragon/config"
	"useless_dragon/setup"
	"useless_dragon/termui"
)

func main() {
	config.Parse()
	player := parsePlayer()
	encounters := setup.ParseEncounters()
	gcui := termui.Create()
	defer gcui.Close()
	for _, enemies := range encounters {
		c := combat.Start(player, enemies...)
		err := termui.RenderCombat(gcui, c)
		if err != nil {
			gcui.Close()
			panic(err)
		}
	}
}
func parsePlayer() *combat.Combatant {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to get cwd")
	}
	data, err := os.ReadFile(fmt.Sprintf("%v/assets/player.json", dir))
	if err != nil {
		log.Fatal("player json not found")
	}
	var p struct {
		Name    string `json:"name"`
		Health  int    `json:"health"`
		Attack  int    `json:"attack"`
		Defence int    `json:"defence"`
	}
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Fatal("failed to unmarshal player json")
	}
	player := combat.Combatant{
		Name:    p.Name,
		Health:  p.Health,
		Attack:  p.Attack,
		Defence: p.Defence,
		Stance:  combat.NormalStance,
		Buffs:   make(map[string]combat.Buff),
	}
	return &player
}
