package combat

import "fmt"

const (
	Won     = 1
	Lost    = 0
	Playing = 2
)
const (
	NormalStance     = "Normal"
	DefensiveStance  = "Defensive"
	AggressiveStance = "Aggressive"
)

type Combat struct {
	Rounds           []Round
	Player           *Combatant
	Enemies          []*Enemy
	Status           int
	PlayerActionChan chan PlayerAction
	UpdateUi         chan bool
}

type Combatant struct {
	Name    string
	Health  int
	Attack  int
	Defence int
	Stance  string
	Buffs   map[string]Buff
}
type Enemy struct {
	Combatant
	Intents []Executable
}
type PlayerAction struct {
	Action Executable
	Target *Enemy
}

func (playerAction *PlayerAction) Execute(round *Round) {
	playerAction.Action.Execute(round.Player, &playerAction.Target.Combatant)
}
func Start(player *Combatant, enemies ...*Enemy) *Combat {
	rounds := make([]Round, 0)
	playerActionChan := make(chan PlayerAction)
	updateUi := make(chan bool)
	combat := &Combat{
		Rounds:           rounds,
		Player:           player,
		Enemies:          enemies,
		Status:           Playing,
		PlayerActionChan: playerActionChan,
		UpdateUi:         updateUi,
	}
	round := NewRound(combat, player, enemies)
	combat.Rounds = append(combat.Rounds, *round)
	go func() {
		for {
			playerAction := <-combat.PlayerActionChan
			combat.AddPlayerAction(&playerAction)
			combat.ResolveRound()
			if combat.Status != Playing {
				close(combat.UpdateUi)
				close(combat.PlayerActionChan)
				break
			} else {
				combat.UpdateUi <- true
			}
			combat.StartNewRound()
		}
	}()
	return combat
}
func (combat *Combat) LastRound() *Round {
	return &combat.Rounds[len(combat.Rounds)-1]
}
func (combat *Combat) ResolveRound() {
	intentIndex := len(combat.Rounds) - 1
	combat.LastRound().ResolveEnemyGuards(intentIndex)
	combat.LastRound().ResolvePlayerAction()
	combat.LastRound().ResolveEnemyIntents(intentIndex)
	combat.LastRound().RoundEnd()
}
func (combat *Combat) StartNewRound() {
	combat.Rounds = append(combat.Rounds, *NewRound(combat, combat.Player, combat.Enemies))
}
func (combat *Combat) AddPlayerAction(action *PlayerAction) {
	combat.Rounds[len(combat.Rounds)-1].PlayerAction = action
}

func (combatant *Combatant) addBuff(buff Buff) {
	existingBuff, ok := combatant.Buffs[buff.Props().Name]
	if ok {
		existingBuff.Props().Duration++
	} else {
		buff.Apply(combatant)
	}
	combatant.Buffs[buff.Props().Name] = buff
}
func (combatant *Combatant) RemoveBuff(buffName string) error {
	buff, ok := combatant.Buffs[buffName]
	if !ok {
		return fmt.Errorf("Buff %s does not exist", buffName)
	}
	buff.Revert(combatant)
	delete(combatant.Buffs, buffName)
	return nil
}
func (combatant *Combatant) prolongBuff(buffName string, duration int) error {
	buff, ok := combatant.Buffs[buffName]
	if !ok {
		return fmt.Errorf("Buff %s does not exist", buffName)
	}
	buff.Props().Duration += duration
	return nil
}
func AddDuration(buff Buff, duration int) {
	buff.Props().Duration += duration
}
func ReduceDuration(buff Buff, duration int) {
	buff.Props().Duration -= duration
}
func RemoveExpiredBuffs(combatants ...*Combatant) {
	for i := range combatants {
		combatant := combatants[i]
		for j := range combatant.Buffs {
			buff := combatant.Buffs[j]
			if buff.Props().Duration <= 0 {
				buff.(Revertable).Revert(combatant)
				switch buff.(type) {
				case *RecklessBuff:
					combatant.Stance = NormalStance
				case *StalwartBuff:
					combatant.Stance = NormalStance
				}
				delete(combatant.Buffs, j)
			}
		}
	}
}
func AreDead(monsters []*Enemy) bool {
	for _, monster := range monsters {
		if monster.Health > 0 {
			return false
		}
	}
	return true
}
func ResolveActionStanceMatrix(combatant *Combatant, action Executable) {
	switch combatant.Stance {
	case NormalStance:
		switch action.(type) {
		case *PowerAttack:
			combatant.Stance = AggressiveStance
			buff := newRecklessBuff()
			combatant.addBuff(buff)
		case *Guard:
			combatant.Stance = DefensiveStance
			combatant.addBuff(newStalwartBuff())
		}
	case DefensiveStance:
		switch action.(type) {
		case *PowerAttack:
			combatant.Stance = NormalStance
			combatant.RemoveBuff(Stalwart)
		case *Guard:
			combatant.prolongBuff(Stalwart, 1)
		}
	case AggressiveStance:
		switch action.(type) {
		case *Guard:
			combatant.Stance = NormalStance
			combatant.RemoveBuff(Reckless)
		case *BasicAttack:
			combatant.prolongBuff(Reckless, 1)
		}
	}
}
func (combat *Combat) EnemyIntentName(enemy *Enemy) string {
	intent := enemy.Intents[AdjustIntentIndex(len(combat.Rounds)-1, len(enemy.Intents))]
	switch action := intent.(type) {
	case *BasicAttack:
		return fmt.Sprintf("Basic Attack: %s is going to deal %v damage ", enemy.Name, action.Calculate(&enemy.Combatant, combat.Player))
	case *PowerAttack:
		return fmt.Sprintf("Power Attack: %s is going to deal %v damage ", enemy.Name, action.Calculate(&enemy.Combatant, combat.Player))
	case *Guard:
		return fmt.Sprintf("Guard: %s is going to gain %v defence ", enemy.Name, action.Calculate(&enemy.Combatant, combat.Player))
	default:
		return "Unknown"
	}
}
