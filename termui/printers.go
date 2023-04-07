package termui

import (
	"fmt"
	"useless_dragon/combat"

	"github.com/jroimartin/gocui"
)

func printStats(v *gocui.View, combatant *combat.Combatant) {
	fmt.Fprintln(v, "Health:", combatant.Health)
	fmt.Fprintln(v, "Attack:", combatant.Attack)
	fmt.Fprintln(v, "Defence:", combatant.Defence)
	fmt.Fprintln(v, "Stance:", combatant.Stance)
}
func printActions(v *gocui.View, player, enemy *combat.Combatant) {
	fmt.Fprintf(v, "→ Basic Attack\ndeal %v damage\n", combat.CreateBasicAttack().Calculate(player, enemy))
	fmt.Fprintf(v, "→ Power Attack\ndeal %v damage\n", combat.CreatePowerAttack().Calculate(player, enemy))
	fmt.Fprintf(v, "→ Guard\ngain %v defence\n", combat.CreateGuard().Calculate(player, enemy))
}
func printBuffs(v *gocui.View, buffs map[string]combat.Buff) {
	fmt.Fprintln(v, "---")
	fmt.Fprintln(v, "Buffs")
	for _, buff := range buffs {
		fmt.Fprintf(v, "%v %v: %v", buff.Props().Name, buff.Props().Duration, buff.Props().Description)
	}
}
