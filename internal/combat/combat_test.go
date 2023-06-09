package combat_test

import (
	"sync"
	"testing"
	"time"

	"github.com/moby-it/useless_dragon/internal/combat"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func setup(t *testing.T) (*combat.Combatant, []*combat.Enemy) {
	t.Helper()
	player := combat.Combatant{
		Name:    "Test_Dragon",
		Health:  100,
		Attack:  30,
		Defence: 10,
		Stance:  "Normal",
		Buffs:   make(map[string]combat.Buff),
	}
	enemies := make([]*combat.Enemy, 0)
	enemies = append(enemies, &combat.Enemy{
		Combatant: combat.Combatant{
			Name:    "Test_Wyvern",
			Health:  30,
			Attack:  10,
			Defence: 5,
			Stance:  "Normal",
			Buffs:   make(map[string]combat.Buff),
		},
		Intents: []combat.Intent{
			combat.CreateBasicAttack(),
			combat.CreatePowerAttack(),
			combat.CreatePowerAttack(),
		},
	})
	return &player, enemies
}
func TestCombatStart(t *testing.T) {
	wg := &sync.WaitGroup{}
	t.Log("When a combat starts")
	player, enemies := setup(t)
	enemyInitialHealth := enemies[0].Health
	playerInitialHealth := player.Health
	c := combat.Start(wg, player, enemies...)
	{
		if c.Player.Health == 100 && len(c.Enemies) > 0 {
			t.Logf("\t Run a smoke Test %v", Success)
		} else {
			t.Fatalf("\t run a smoke Test %v", Failed)
		}
		firstEnemy := enemies[0]
		playerAttackDamage := combat.CreateBasicAttack().Calculate(c.Player, &firstEnemy.Combatant)
		enemyAttackDamage := firstEnemy.Intents[0].Calculate(&firstEnemy.Combatant, player)
		t.Log("When the player does a basic attack")
		{
			action := combat.PlayerAction{
				Action: combat.CreateBasicAttack(),
				Target: enemies[0],
			}
			c.PlayerActionChan <- action
			select {
			case <-c.UpdateUi:
				t.Logf("\t UpdateUi channel should receive a value %v", Success)
			case <-time.After(1 * time.Second):
				t.Fatalf("\t UpdateUi channel should receive a value %v", Failed)
			}
			if firstEnemy.Health == enemyInitialHealth-playerAttackDamage {
				t.Logf("\t Enemy remaining Health should equal Starting Health - Attack Damage %v", Success)
			} else {
				t.Logf("\t Enemy remaining Health should equal Starting Health - Attack Damage %v", Failed)
			}
		}
		t.Log("After the round is resolved")
		{
			if c.Player.Health == playerInitialHealth-enemyAttackDamage {
				t.Logf("\t Player remaining Health should equal Starting Health - Enemy First Intent Damage %v", Success)
			} else {
				t.Logf("\t Player remaining Health should equal Starting Health - Enemy First Intent Damage %v", Failed)
			}
		}
	}
}
