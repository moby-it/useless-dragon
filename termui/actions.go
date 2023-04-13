package termui

import (
	"fmt"
	"useless_dragon/combat"

	"github.com/jroimartin/gocui"
)

const ActionsWidth = 40
const ActionsHeight = 10

var selectedAction combat.Executable

func CreateActions(g *gocui.Gui, c *combat.Combat) error {
	const topPadding = 18
	const leftPadding = 1
	player := c.Player
	enemy := c.Enemies[0]
	if v, err := g.SetView(ActionsView, leftPadding, topPadding, leftPadding+ActionsWidth, topPadding+ActionsHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Actions"
		v.Wrap = true
		printActions(v, player, &enemy.Combatant)
	}
	return nil
}
func UpdateActions(g *gocui.Gui, c *combat.Combat) error {
	player := c.Player
	enemy := c.Enemies[0]
	v, err := g.View(ActionsView)
	if err != nil {
		return err
	}
	v.Clear()
	selectedAction = nil
	g.SetCurrentView(ActionsView)
	printActions(v, player, &enemy.Combatant)
	return nil
}
func selectAction(c *combat.Combat) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if c.Player.Health <= 0 {
			return nil
		}
		_, cy := v.Cursor()
		if cy == 0 || cy == 1 {
			selectedAction = combat.CreateBasicAttack()
			// playerAction = combat.PlayerAction{Action: combat.CreateBasicAttack(), Target: enemy}
			// c.PlayerActionChan <- playerAction

		} else if cy == 2 || cy == 3 {
			selectedAction = combat.CreatePowerAttack()
		} else if cy == 4 || cy == 5 {
			selectedAction = combat.CreateGuard()
		}
		g.SetCurrentView(fmt.Sprintf("%v_0", EnemyStatsView))
		return nil
	}
}
func cancelAction(c *combat.Combat) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		g.SetCurrentView(ActionsView)
		selectedAction = nil
		return nil
	}
}
