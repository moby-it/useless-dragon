package termui

import (
	"fmt"
	"strings"

	"github.com/moby-it/useless_dragon/internal/config"
)

func renderActions(m model) string {
	player := m.combat.Player
	title := headerStyle.Render("Actions")
	ba := fmt.Sprintf("\n  Basic Attack: %v damage", player.Attack)
	pa := fmt.Sprintf("\n  Power Attack: %v damage", player.Attack+config.Get().PowerAttackBonus)
	gd := fmt.Sprintf("\n  Guard: Gain %v Defense", config.Get().Guard_Bonus)
	if m.selectedAction < 0 {
		switch m.cursor {
		case 0:
			ba = addActiveCursorPointer(ba)
		case 1:
			pa = addActiveCursorPointer(pa)
		case 2:
			gd = addActiveCursorPointer(gd)
		}
	} else {
		switch m.selectedAction {
		case 0:
			ba = addActiveCursorPointer(ba)
		case 1:
			pa = addActiveCursorPointer(pa)
		case 2:
			gd = addActiveCursorPointer(gd)
		}
	}
	actions := title + ba + pa + gd
	if m.selectedAction < 0 {
		actions = focusedBoxStyle.Render(actionsStyle.Render(actions))
	} else {
		actions = boxStyle.Render(actionsStyle.Render(actions))
	}
	return actions
}
func addActiveCursorPointer(text string) string {
	return hoverSyle.Render(strings.Replace(text, " ", ">", 1))
}
