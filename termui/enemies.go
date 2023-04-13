package termui

import (
	"fmt"
	"strconv"
	"strings"
	"useless_dragon/combat"

	"github.com/jroimartin/gocui"
)

const boxWidth = 65
const paddingLeft = 45

func createEnemyStats(g *gocui.Gui, c *combat.Combat, idx int) error {
	enemy := c.Enemies[idx]
	padding := paddingLeft + (boxWidth * idx)
	if idx > 0 {
		padding += 5
	}
	if v, err := g.SetView(fmt.Sprintf("%v_%v", EnemyStatsView, idx), padding, 3, padding+boxWidth, 15); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = enemy.Name
		v.Wrap = true
		fmt.Fprint(v, "Intent:\n", c.EnemyIntentName(enemy), "\n")
		fmt.Fprintln(v, "---")
		printStats(v, &enemy.Combatant)
		if len(enemy.Buffs) > 0 {
			printBuffs(v, enemy.Buffs)
			fmt.Fprintln(v, "---")

		}
	}
	return nil
}
func updateEnemyStats(g *gocui.Gui, c *combat.Combat, idx int) error {
	enemy := c.Enemies[idx]

	v, err := g.View(fmt.Sprintf("%v_%v", EnemyStatsView, idx))
	if err != nil {
		return err
	}
	v.Clear()
	if enemy.Health <= 0 {
		fmt.Fprintln(v, "You won! :)")
		return nil
	}
	fmt.Fprint(v, "Intent:\n", c.EnemyIntentName(enemy), "\n")
	fmt.Fprintln(v, "---")
	printStats(v, &enemy.Combatant)
	if len(enemy.Buffs) > 0 {
		printBuffs(v, enemy.Buffs)
	}
	return nil
}
func selectEnemy(c *combat.Combat) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if strings.Contains(v.Name(), EnemyStatsView) {
			idx := strings.Split(v.Name(), "_")[1]
			i, err := strconv.Atoi(idx)
			if err != nil {
				return err
			}
			enemy := c.Enemies[i]
			if enemy.Health < 0 {
				return nil
			}
			playerAction := combat.PlayerAction{
				Action: selectedAction,
				Target: c.Enemies[i],
			}
			c.PlayerActionChan <- playerAction
			return nil
		}
		return nil
	}
}
func nextEnemy(c *combat.Combat) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if strings.Contains(v.Name(), EnemyStatsView) {
			idx := strings.Split(v.Name(), "_")[1]
			i, err := strconv.Atoi(idx)
			if err != nil {
				return err
			}
			if i < len(c.Enemies) {
				g.SetCurrentView(fmt.Sprintf("%v_%v", EnemyStatsView, i+1))
			}
			return nil
		}
		return nil
	}
}
func previousEnemy(c *combat.Combat) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if strings.Contains(v.Name(), EnemyStatsView) {
			idx := strings.Split(v.Name(), "_")[1]
			i, err := strconv.Atoi(idx)
			if err != nil {
				return err
			}
			if i > 0 {
				g.SetCurrentView(fmt.Sprintf("%v_%v", EnemyStatsView, i-1))
			}
			return nil
		}
		return nil
	}
}
