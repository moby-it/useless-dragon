# Brief Game Combat Summary

Made with a friend called gpt4

# Game Summary

## 1. Stances

Stances are combat states that affect the player's abilities and interactions during turn-based combat. There are three primary stances: Normal (N), Aggressive (AS), and Defensive (DS). Stances have round limits, encouraging players to strategically switch between them.

## 2. Actions

Players can perform various actions during combat, such as Basic Attack, Power Strike, Guard, and Taunt. The effects of these actions can change depending on the player's current stance.

## 3. Stance-Action Matrix

The Stance-Action Matrix shows the relationship between actions and the resulting stance changes. It serves as a blueprint for the dynamic combat system.

| Current Status | Basic Attack | Power Strike | Guard   | Taunt  |
|----------------|--------------|--------------|---------|--------|
| Normal (N)     | N            | AS (3 turns) | DS (3 turns) | N      |
| Aggressive (AS)| N            | N            | DS (3 turns) | N      |
| Defensive (DS) | N            | ST (Opponent)| DS (3 turns) | N      |
| Stunned (ST)   | -            | -            | -       | -      |

## 4. Skills

Players can acquire skills through leveling up and investing in a mastery tree. The mastery tree is divided into Aggressive, Defensive, and Normal branches, allowing players to customize their playstyle.

## 5. Items

Lootable gear and potions can provide bonuses or penalties to specific stances, further enhancing character progression and customization. Players can collect and equip different items to optimize their combat abilities.

### Example Lists

#### Unique Heritage Skills

1. **Dragon's Breath**: Breathe fire in Aggressive Stance after 3 rounds, dealing significant damage and stunning the player for the next turn.
2. **Unyielding Defender**: Negate all damage from an incoming attack when in Defensive Stance for more than 3 rounds.
3. **Arcane Channeling**: Increased effectiveness for elemental spells and abilities when in Normal Stance.
4. **Vampiric Strikes**: Convert a percentage of damage dealt by Basic Attacks and Power Strikes into Health when in Aggressive Stance.
5. **Guardian's Aura**: Provide damage reduction to nearby allies when in Defensive Stance.
6. **Quick Reflexes**: Gain a chance to dodge incoming attacks when in Normal Stance.

#### Stance-Affecting Items

1. **Berserker's Axe**: Increases damage dealt in Aggressive Stance but reduces defense in Defensive Stance.
2. **Shield of Fortitude**: Increases damage reduction in Defensive Stance but reduces damage dealt in Aggressive Stance.
3. **Balanced Blade**: Slightly increases damage dealt in all stances.
4. **Boots of Agility**: Increases the chance to dodge attacks in Normal Stance.
5. **Helm of Focus**: Extends the duration of Aggressive Stance by one turn.
6. **Chestplate of Resilience**: Extends the duration of Defensive Stance by one turn.
7. **Ring of Adaptability**: Reduces the number of turns required to switch between stances.
8. **Amulet of the Tactician**: Increases the chance to gain an extra action in Normal Stance.
