package combat

import "github.com/moby-it/useless_dragon/internal/config"

type Executable interface {
	Execute(initiator, receiver *Combatant)
}
type Calculatable interface {
	Calculate(initiator *Combatant, receiver *Combatant) int
}
type Intent interface {
	Executable
	Calculatable
}

// Basic Attack

type BasicAttack struct {
	Name        string
	Description string
}

func (a BasicAttack) Calculate(initiator, receiver *Combatant) int {
	damage := initiator.Attack - receiver.Defence
	if damage < 0 {
		return 0
	}
	return damage
}
func CreateBasicAttack() *BasicAttack {
	return &BasicAttack{
		Name:        "Basic Attack",
		Description: "A basic attack",
	}
}
func (a BasicAttack) Execute(initiator *Combatant, receiver *Combatant) {
	damage := a.Calculate(initiator, receiver)
	receiver.Health -= damage
}

// Power Attack

type PowerAttack struct {
	Name        string
	Description string
}

func (a PowerAttack) Calculate(initiator *Combatant, receiver *Combatant) int {
	damage := initiator.Attack + config.Get().PowerAttackBonus - receiver.Defence
	if damage < 0 {
		return 0
	}
	return damage
}

func CreatePowerAttack() *PowerAttack {
	return &PowerAttack{
		Name:        "Power Attack",
		Description: "A powerful attack",
	}
}

func (a PowerAttack) Execute(initiator *Combatant, receiver *Combatant) {
	damage := a.Calculate(initiator, receiver)
	receiver.Health -= damage
}

// Guard

type Guard struct {
	Name        string
	Description string
}

func CreateGuard() *Guard {
	return &Guard{
		Name:        "Guard",
		Description: "Guard against incoming attacks",
	}
}

func (a Guard) Execute(initiator, receiver *Combatant) {
	defence := a.Calculate(initiator, nil)
	buff := newFortifiedBuff(defence)
	initiator.addBuff(buff)
}
func (a Guard) Calculate(initiator, receiver *Combatant) int {
	return config.Get().Guard_Bonus
}
