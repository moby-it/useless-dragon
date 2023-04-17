package main

import (
	"fmt"
	"os"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/moby-it/useless_dragon/cmd/termui"
	"github.com/moby-it/useless_dragon/internal/combat"
	"github.com/moby-it/useless_dragon/internal/config"
	"github.com/moby-it/useless_dragon/internal/setup"
)

func main() {
	wg := &sync.WaitGroup{}
	// explicitly initialy game config
	config.Parse()
	player := setup.ParsePlayer()
	encounters := setup.ParseEncounters()
	wg.Add(len(encounters))
	updaterChan := make(chan *combat.Combat)
	p := tea.NewProgram(termui.InitialModel(updaterChan), tea.WithAltScreen())
	go func(wg *sync.WaitGroup) {
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	}(wg)
	for _, enemies := range encounters {
		c := combat.Start(wg, player, enemies...)
		updaterChan <- c
	}
	wg.Wait()
}
