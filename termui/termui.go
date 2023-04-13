package termui

import (
	"fmt"
	"useless_dragon/combat"

	"github.com/jroimartin/gocui"
)

const (
	ActionsView     = "actions"
	PlayerStatsView = "playerStats"
	EnemyStatsView  = "enemyStats"
)

func Create() *gocui.Gui {
	// if cannot create cui panic
	g, _ := gocui.NewGui(gocui.Output256)
	g.Cursor = true
	g.Highlight = true
	g.InputEsc = true
	g.SelFgColor = gocui.ColorBlue
	return g
}
func RenderCombat(g *gocui.Gui, c *combat.Combat) error {
	g.SetManagerFunc(combatLayout(c))
	registerCombatKeybindings(g, c)
	go func(g *gocui.Gui, c *combat.Combat) {
		for range c.UpdateUi {
			UpdateUI(g, c)
		}
	}(g, c)
	err := g.MainLoop()
	if err == gocui.ErrQuit && c.Status == combat.Playing {
		return fmt.Errorf("game crashed")
	}
	return nil
}

func combatLayout(c *combat.Combat) func(*gocui.Gui) error {
	return func(g *gocui.Gui) error {
		createActions(g, c)
		createPlayerStats(g, c)
		for i := range c.Enemies {
			err := createEnemyStats(g, c, i)
			if err != nil {
				return err
			}
		}
		if g.CurrentView() == nil {
			g.SetCurrentView(ActionsView)
		}
		if c.Status == combat.Won || c.Status == combat.Lost {
			return gocui.ErrQuit
		}
		return nil
	}

}
func registerCombatKeybindings(g *gocui.Gui, c *combat.Combat) {
	// keybidings
	g.SetKeybinding(ActionsView, gocui.KeyCtrlC, gocui.ModNone, quit)
	g.SetKeybinding(ActionsView, gocui.KeyArrowDown, gocui.ModNone, cursorDown)
	g.SetKeybinding(ActionsView, gocui.KeyArrowUp, gocui.ModNone, cursorUp)
	g.SetKeybinding(ActionsView, gocui.KeyEnter, gocui.ModNone, selectAction(c))
	g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, selectEnemy(c))
	g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, nextEnemy(c))
	g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, previousEnemy(c))
	g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, cancelAction(c))
}
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
func cursorDown(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	if cy < 5 {
		if err := v.SetCursor(0, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
	}
	return nil
}
func cursorUp(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	if cy > 0 {
		if err := v.SetCursor(0, cy-1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}
func UpdateUI(cui *gocui.Gui, c *combat.Combat) {
	cui.Update(func(g *gocui.Gui) error {
		// update actions
		updateActions(g, c)
		// update enemies
		for i := range c.Enemies {
			updateEnemyStats(g, c, i)
		}
		// update player stats
		updatePlayerStats(g, c)
		return nil
	})
}
