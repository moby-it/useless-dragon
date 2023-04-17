package setup

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/moby-it/useless_dragon/internal/combat"
)

func ParsePlayer() *combat.Combatant {
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
