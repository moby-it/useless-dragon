package setup

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moby-it/useless_dragon/internal/combat"
)

type EnemyJSON struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Health      int      `json:"health"`
	Attack      int      `json:"attack"`
	Defense     int      `json:"defense"`
	Stance      string   `json:"stance"`
	Intents     []string `json:"intents"`
}

func ParseEncounters() [][]*combat.Enemy {
	// get the current directory
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	encountersFilepath := filepath.Join(dir, "assets/encounters.json")
	var data []byte
	if data, err = os.ReadFile(encountersFilepath); err != nil {
		panic(err)
	}
	var encounters [][]string
	err = json.Unmarshal(data, &encounters)
	if err != nil {
		panic(err)
	}
	enemies := make([][]*combat.Enemy, len(encounters))

	var enemyJSON EnemyJSON
	for i, encounter := range encounters {
		e := make([]*combat.Enemy, len(encounter))
		for j, enemyName := range encounter {
			filepath := filepath.Join(dir, "assets/enemies", fmt.Sprintf("%v.json", enemyName))
			data, err := os.ReadFile(filepath)
			if err != nil {
				panic(err)
			}
			json.Unmarshal(data, &enemyJSON)
			e[j] = transformMonsterJSON(&enemyJSON)
		}
		enemies[i] = e
	}
	return enemies
}
func transformMonsterJSON(monster *EnemyJSON) *combat.Enemy {
	// create a new monster
	m := &combat.Enemy{
		Combatant: combat.Combatant{
			Name:    monster.Name,
			Health:  monster.Health,
			Attack:  monster.Attack,
			Defence: monster.Defense,
			Stance:  monster.Stance,
			Buffs:   make(map[string]combat.Buff),
		},
	}
	// loop through the intents
	for _, intent := range monster.Intents {
		var action combat.Executable
		switch intent {
		case "ba":
			action = combat.CreateBasicAttack()
		case "pa":
			action = combat.CreatePowerAttack()
		case "gd":
			action = combat.CreateGuard()
		}
		// append the intent to the monster
		switch a := action.(type) {
		case *combat.BasicAttack:
			m.Intents = append(m.Intents, a)
		case *combat.PowerAttack:
			m.Intents = append(m.Intents, a)
		case *combat.Guard:
			m.Intents = append(m.Intents, a)
		}
	}
	// return the monster
	return m
}
