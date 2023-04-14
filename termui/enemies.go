package termui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func renderEnemies(m model) string {
	var enemyText string
	enemieViews := make([]string, len(m.combat.Enemies))
	for i, enemy := range m.combat.Enemies {
		title := headerStyle.Render(enemy.Name)
		if enemy.Health <= 0 {
			enemyText = headerStyle.Render(enemy.Name + " is dead")
		} else {
			intent := "\n" + m.combat.EnemyIntentName(enemy)
			stats := "\n" + renderStats(&enemy.Combatant, "")
			enemyText = fmt.Sprint(title, intent, stats)
		}
		if m.selectedAction >= 0 && m.cursor == i {
			enemieViews[i] = focusedBoxStyle.Width(enemyBoxWidth).Height(boxHeight).Render(enemyText)
		} else {
			enemieViews[i] = boxStyle.Width(enemyBoxWidth).Height(boxHeight).Render(enemyText)
		}
	}
	return lipgloss.JoinVertical(lipgloss.Bottom, enemieViews...)
}
