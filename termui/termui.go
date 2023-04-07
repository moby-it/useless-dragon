package termui

import (
	"log"
	"useless_dragon/combat"

	"github.com/jroimartin/gocui"
)

const (
	ActionsView     = "actions"
	PlayerStatsView = "playerStats"
	EnemyStatsView  = "enemyStats"
)

func Render(c *combat.Combat) (*gocui.Gui, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}
	g.Cursor = true
	g.Highlight = true
	g.SelFgColor = gocui.ColorBlue
	g.SetManagerFunc(layout(c))
	// keybidings
	if err := g.SetKeybinding(ActionsView, gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return nil, err
	}
	if err := g.SetKeybinding(ActionsView, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return nil, err
	}
	if err := g.SetKeybinding(ActionsView, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return nil, err
	}
	if err := g.SetKeybinding(ActionsView, gocui.KeyEnter, gocui.ModNone, selectAction(c)); err != nil {
		return nil, err
	}
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, selectEnemy(c)); err != nil {
		return nil, err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, nextEnemy(c)); err != nil {
		return nil, err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, previousEnemy(c)); err != nil {
		return nil, err
	}
	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, cancelAction(c)); err != nil {
		return nil, err
	}
	go func(g *gocui.Gui, c *combat.Combat) {
		for range c.UpdateUi {
			UpdateUI(g, c)
		}
	}(g, c)
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	return g, nil
}

func layout(c *combat.Combat) func(*gocui.Gui) error {
	return func(g *gocui.Gui) error {
		CreateActions(g, c)
		CreatePlayerStats(g, c)
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
		UpdateActions(g, c)
		// update enemies
		for i := range c.Enemies {
			updateEnemyStats(g, c, i)
		}
		// update player stats
		UpdatePlayerStats(g, c)
		return nil
	})
}
