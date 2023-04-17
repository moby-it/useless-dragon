package combat

import (
	"fmt"
	"log"
	"strings"

	"github.com/moby-it/useless_dragon/internal/config"
)

const (
	Fortified = "Fortified"
	Reckless  = "Reckless"
	Stalwart  = "Stalwart"
)

type Applicable interface {
	Apply(combatant *Combatant)
}
type Revertable interface {
	Revert(combatant *Combatant)
}

// BuffProperties is a temporary effect that can be applied to a combatant. It can be either positive or negative.
type BuffProperties struct {
	Name        string
	Description string
	Duration    int
}
type Buff interface {
	Props() *BuffProperties
	Applicable
	Revertable
}

// Increaces player defence.
type FortifiedBuff struct {
	BuffProperties
	Defence int
}

func newFortifiedBuff(defence int) *FortifiedBuff {
	return &FortifiedBuff{
		BuffProperties: BuffProperties{
			Name:        "Fortified",
			Description: fmt.Sprintf("+%v def", defence),
			Duration:    1,
		},
		Defence: defence,
	}
}
func (buff *FortifiedBuff) Props() *BuffProperties {
	return &buff.BuffProperties
}
func (buff FortifiedBuff) Apply(combatant *Combatant) {
	combatant.Defence += buff.Defence
}
func (buff FortifiedBuff) Revert(combatant *Combatant) {
	combatant.Defence -= buff.Defence
}

// Aggressive stance buff: Decreases player defence and increases attack.
type RecklessBuff struct {
	BuffProperties
	Attack  int
	Defence int
}

func newRecklessBuff() *RecklessBuff {
	buff, found := config.Get().Buffs["reckless"]
	if !found {
		log.Fatalln("Reckless sbuff not found")
	}
	description := formatDescription(buff)
	return &RecklessBuff{
		BuffProperties: BuffProperties{
			Name:        "Reckless",
			Description: description,
			Duration:    buff.Duration,
		},
		Attack:  buff.Attack,
		Defence: buff.Defence,
	}
}
func (buff *RecklessBuff) Props() *BuffProperties {
	return &buff.BuffProperties
}
func (buff RecklessBuff) Apply(combatant *Combatant) {
	combatant.Attack += buff.Attack
	combatant.Defence += buff.Defence
}
func (buff RecklessBuff) Revert(combatant *Combatant) {
	combatant.Attack -= buff.Attack
	combatant.Defence -= buff.Defence
}

// Deffensive stance buff. Increases player defence.
type StalwartBuff struct {
	BuffProperties
	Defence int
	Attack  int
}

func newStalwartBuff() *StalwartBuff {
	buff, found := config.Get().Buffs["stalwart"]
	if !found {
		log.Fatal("Stalward buff not found")
	}
	return &StalwartBuff{
		BuffProperties: BuffProperties{
			Name:        "Stalwart",
			Description: formatDescription(buff),
			Duration:    buff.Duration,
		},
		Defence: buff.Defence,
		Attack:  buff.Attack,
	}
}

func (buff *StalwartBuff) Props() *BuffProperties {
	return &buff.BuffProperties
}
func (buff StalwartBuff) Apply(combatant *Combatant) {
	combatant.Defence += buff.Defence
	combatant.Attack += buff.Attack
}
func (buff StalwartBuff) Revert(combatant *Combatant) {
	combatant.Defence -= buff.Defence
	combatant.Attack -= buff.Attack
}
func formatDescription(buff config.Buff) string {
	var atk, def string
	if buff.Attack > 0 {
		atk = fmt.Sprintf("+%v atk", buff.Attack)
	} else if buff.Attack < 0 {
		atk = fmt.Sprintf("%v atk", buff.Attack)
	}
	if buff.Defence > 0 {
		def = fmt.Sprintf("+%v def", buff.Defence)
	} else if buff.Defence < 0 {
		def = fmt.Sprintf("%v def", buff.Defence)
	}
	return strings.Join([]string{atk, def}, ",")
}
