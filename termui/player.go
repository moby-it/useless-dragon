package termui

import (
	"fmt"
	"useless_dragon/combat"

	"github.com/jroimartin/gocui"
)

const PlayerStatsWidth = 40
const leftPadding = 1

func createPlayerStats(g *gocui.Gui, c *combat.Combat) error {
	player := c.Player
	v, err := g.SetView(PlayerStatsView, leftPadding, 3, leftPadding+PlayerStatsWidth, 15)
	if err != gocui.ErrUnknownView {
		return err
	}
	v.Title = player.Name
	v.Wrap = true
	printStats(v, player)
	if len(player.Buffs) > 0 {
		printBuffs(v, player.Buffs)
	}
	return nil
}
func updatePlayerStats(g *gocui.Gui, c *combat.Combat) error {
	player := c.Player
	v, err := g.View(PlayerStatsView)
	if err != nil {
		return err
	}
	v.Clear()
	if player.Health <= 0 {
		fmt.Fprintln(v, "You died! :(")
		return nil
	}
	printStats(v, player)
	if len(player.Buffs) > 0 {
		printBuffs(v, player.Buffs)
	}
	return nil
}
