# USELESS DRAGON
some readme changes
## Introduction

Welcome to Useless Dragon, a terminal-based roguelike game where you play as a dragon brought up in a society of dragons who can transform into humans.

Our hero is the Useless Dragon. Dragons have the ability to transform into humans at will and live in a developed society, hunting as dragons when they feel like it. However, Useless Dragon has never been able to transform into a dragon like the rest of his family and friends, leading him to be mocked and called "useless". Little does anyone know, Useless Dragon is actually a Phoenix, a legendary species of dragon that develops their powers much more slowly but has the ability to never truly die.

In our roguelike game, Useless Dragon sets out on an adventure to trigger his dormant dragon abilities and prove to everyone that he is not a useless dragon after all.

*The current state of the game is not even considered alpha. We paused development and brainstoring for the moment and prioritized other things. We resume development at some point in 2024*


# Combat

## Stances

The game features three different stances that the player can take during combat:

1. Human (NS): The default stance.
2. Drakonid (AS): An aggressive stance that focuses on dealing damage.
3. Phoenix (DS): A defensive stance that focuses on surviving and healing.

Each stance has a maximum duration of 3 turns initially (except for Human), after which the player will automatically revert to Human stance.

## Actions

The player has access to three different actions during combat:

1. Basic Attack (BA): A standard attack that deals damage.
2. Power Attack(PA): A stronger attack that deals more damage, but may have additional effects depending on the current stance.
3. Guard(GA): A defensive move that reduces incoming damage and may have additional effects depending on the current stance.

Each action has a different effect depending on the current stance. For example, using a Power Attack in Drakonid (AS) stance deals more damage than using it in another stance.

## Actions-Stances Matrix

|              | Normal (NS) | Aggressive (AS) | Defensive (DS) |
|--------------|------------|---------------|---------------|
| Basic Attack |       ✅     |        ✅ <br>stance duration +1|        ✅      |
| Power Attack | Aggressive (AS)| Aggressive (AS)   | Normal (NS)     |
| Guard        | Defensive (DS) | Normal (NS)     | Defensive (DS)<br>stance duration +1|
| Max Duration |       ∞     |         3       |         3       |

## Progression

Useless Dragon progresses in two ways. With **leveling up** and **Items**.

### Leveling Up

As the player defeats monsters, they gain experience points (XP) that allow them to level up. Each level-up rewards the player with Mastery Points (MP) that they can invest in one of three skill trees: Aggressive, Defensive, and Normal.

Here is how some example skills could look like:

 Name | Description 
------------|--------------- 
Dragon's Breath	| Breathe fire in Aggressive Stance after 3 rounds, dealing significant damage and stunning the player for the next turn.
 Unyielding Defender | Negate all damage from an incoming attack when in Defensive Stance for more than 3 rounds.  
Arcane Channeling | Increased effectiveness for elemental spells and abilities when in Normal Stance.
Vampiric Strikes	| Convert a percentage of damage dealt by Basic Attacks and Power Strikes into Health when in Aggressive Stance.
Guardian's Aura	 | Provide damage reduction to nearby allies when in Defensive Stance.


## Items

In this game monsters will drop **Items** regarding what they carry when they fight you. Items might also be stance related. 

Here are how some example items might look like:

 Name | Description 
------------|--------------- 
Berserker's Axe		| Increases damage dealt in Aggressive Stance but reduces defense in Defensive Stance.
 Shield of Fortitude	 | Increases damage reduction in Defensive Stance but reduces damage dealt in Aggressive Stance.
Balanced Blade	 | Slightly increases damage dealt in all stances.
Helm of Focus	| Extends the duration of Aggressive Stance by one turn.
Chestplate of Resilience	 | Extends the duration of Defensive Stance by one turn.
Ring of Adaptability | Reduces the number of turns required to switch between stances.
Amulet of the Tactician	| Increases the chance to gain an extra action in Normal Stance.

# Monsters

Monsters obey the same ruleset regarding Stances and Actions. The player can see the action which the monster will conduct as well as the current stance. Each monster can have it's own unique Skills that affect Actions/Stance in a unique way. 

We have one example monster for now.

## Teenage Wyvern

The Teenage Wyvern is a young, arrogant dragon that always mocked the Useless Dragon. During the Useless Dragon's first adventure, the Wyvern attacks him, forcing him to fight. The Wyvern prefers the Drakonid (AS) stance and never switches to the Phoenix (DS) stance. However, the Wyvern receives 10% more damage each turn spent in Drakonid (AS) stance.
