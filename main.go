package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"useless_dragon/combat"
	"useless_dragon/config"
	"useless_dragon/setup"
	"useless_dragon/termui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	wg := &sync.WaitGroup{}
	config.Parse()
	player := parsePlayer()
	encounters := setup.ParseEncounters()
	wg.Add(len(encounters))
	updaterChan := make(chan *combat.Combat)
	p := tea.NewProgram(termui.InitialModel(updaterChan))
	go func(wg *sync.WaitGroup) {
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}(wg)
	for i, enemies := range encounters {
		c := combat.Start(wg, player, enemies...)
		log.Println("to start combat", i)
		updaterChan <- c
		log.Println("combat", i, "ended")
	}
	wg.Wait()
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
