package combat

const (
	ChoosingAction = iota + 1
	ResolveAction
	ResolveIntents
	RoundEnd
)

type Round struct {
	Status       int
	PlayerAction *PlayerAction
	Player       *Combatant
	Enemies      []*Enemy
	Combat       *Combat
}

func NewRound(combat *Combat, player *Combatant, enemies []*Enemy) *Round {
	return &Round{
		Status:       ChoosingAction,
		Player:       player,
		Enemies:      enemies,
		PlayerAction: nil,
		Combat:       combat,
	}
}
func (round *Round) ResolveEnemyGuards(intentIndex int) {
	for _, enemy := range round.Enemies {
		intentIndex = AdjustIntentIndex(intentIndex, len(enemy.Intents))
		intent := enemy.Intents[intentIndex]
		switch intent.(type) {
		case *Guard:
			intent.Execute(&enemy.Combatant, nil)
		}
	}
}
func (round *Round) ResolvePlayerAction() {
	round.PlayerAction.Execute(round)
	won := true
	for _, enemy := range round.Enemies {
		if enemy.Health > 0 {
			won = false
			break
		}
	}
	if won {
		round.Combat.Status = Won
	}
}
func (round *Round) ResolveEnemyIntents(intentIndex int) {
	if AreDead(round.Enemies) {
		return
	}
	for _, enemy := range round.Enemies {
		intentIndex = AdjustIntentIndex(intentIndex, len(enemy.Intents))
		intent := enemy.Intents[intentIndex]
		switch intent.(type) {
		case *Guard:
			// Do nothing. Guard is resolved in ResolveEnemyGuards
		default:
			intent.Execute(&enemy.Combatant, round.Player)
		}
	}
	lost := round.Player.Health <= 0
	if lost {
		round.Combat.Status = Lost
	}
	// Resolve the intents
}
func (round *Round) RoundEnd() {
	// Reduce Buff Durations
	for i := range round.Player.Buffs {
		buff := round.Player.Buffs[i]
		ReduceDuration(buff, 1)
	}

	for i := range round.Enemies {
		enemy := round.Enemies[i]
		for i := range enemy.Buffs {
			buff := enemy.Buffs[i]
			ReduceDuration(buff, 1)
		}
	}
	// Resolve Stance Action Matrix
	ResolveActionStanceMatrix(round.Player, round.Combat.LastRound().PlayerAction.Action)
	for _, enemy := range round.Enemies {
		intentIndex := AdjustIntentIndex(len(round.Combat.Rounds)-1, len(enemy.Intents))
		intent := enemy.Intents[intentIndex]
		ResolveActionStanceMatrix(&enemy.Combatant, intent)
	}

	// Remove expired buffs
	combatants := make([]*Combatant, 0)
	combatants = append(combatants, round.Player)
	for _, enemy := range round.Enemies {
		combatants = append(combatants, &enemy.Combatant)
	}
	RemoveExpiredBuffs(combatants...)
}

// If the intent index is greater than the number of intents, it will loop back around to the beginning
func AdjustIntentIndex(intentIndex, intentsLength int) int {
	if intentIndex >= intentsLength {
		intentIndex = intentIndex % intentsLength
	}
	return intentIndex
}
