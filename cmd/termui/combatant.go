package termui

import (
	"fmt"

	"github.com/moby-it/useless_dragon/internal/combat"
)

func renderStats(combatant *combat.Combatant, title string) string {
	if combatant.Health <= 0 {
		return headerStyle.Render(fmt.Sprintf("%v is dead", combatant.Name))
	}
	stats := fmt.Sprintf("\nHealth: %v\nAttack: %v\nDefense: %v", combatant.Health, combatant.Attack, combatant.Defence)
	if len(title) > 0 {
		stats = title + stats
	}
	buffs, hasBuffs := renderBuffs(combatant)
	if hasBuffs {
		stats += buffs
	}
	return statsStyle.Render(stats)
}
func renderBuffs(combatant *combat.Combatant) (string, bool) {
	if len(combatant.Buffs) > 0 {
		s := bold.Render("\nBuffs")
		for _, buff := range combatant.Buffs {
			s += fmt.Sprintf("\n%v %v: %v", buff.Props().Name, buff.Props().Duration, buff.Props().Description)
		}
		return s, true
	}
	return "", false
}
func renderCombatant(combatant *combat.Combatant) string {
	title := headerStyle.Render(combatant.Name)
	if combatant.Health <= 0 {
		return statsStyle.Render(combatant.Name + " is dead")
	}
	s := renderStats(combatant, title)
	return s
}
