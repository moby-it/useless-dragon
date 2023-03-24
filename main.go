package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	name       string
	hp         int
	maxHP      int
	attack     int
	defense    int
	level      int
	experience int
	mana       int
}

type Enemy struct {
	name   string
	hp     int
	attack int
}

func main() {
	// Initialize the player and enemy
	player := Player{name: "Player", hp: 100, maxHP: 100, attack: 10, defense: 5, level: 1, experience: 0, mana: 100}
	enemy := Enemy{name: "Dragon", hp: 50, attack: 5}

	fmt.Println("Choose your Champion!")
	fmt.Println("1. Fighter")
	fmt.Println("2. Ranger")
	var champion int
	fmt.Scanln(&champion)
	if champion == 1 {
		player.hp += 100
		player.defense += 5
		fmt.Println("You are a Fighter!")
	} else {
		player.attack += 10
		fmt.Println("You are a Ranger!")
	}
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Main game loop
	for {
		fmt.Println("Player HP:", player.hp)
		fmt.Println("Enemy HP:", enemy.hp)

		// Player turn
		fmt.Println("Player's turn")
		fmt.Println("1. Attack")
		fmt.Println("2. Defend")
		fmt.Println("3. Cast Spell")
		fmt.Println("4. Run Away")
		var choice int
		fmt.Scanln(&choice)
		if choice == 1 {
			damage := player.attack - enemy.attack
			if damage < 0 {
				damage = 0
			}
			enemy.hp -= damage
			fmt.Println("Player attacks for", damage, "damage")
		} else if choice == 3 {

			damage := 20

			enemy.hp -= damage
			fmt.Println("You cast a fireball!")
			fmt.Println("Player attacks for", damage, "damage")
		} else if choice == 4 {

			fmt.Println("Run Away Safely!")
			break
		} else {
			fmt.Println("Player defends")
			player.defense += 5
		}

		// Check if enemy is defeated
		if enemy.hp <= 0 {
			fmt.Println("Player defeats", enemy.name)
			fmt.Println("Congratulations!", "You are the best!")
			player.experience += 20
			if player.experience >= 100 {
				player.level++
				player.maxHP += 10
				player.attack += 2
				player.defense += 2
				player.hp = player.maxHP
				player.experience = 0
				fmt.Println("Player leveled up to level", player.level)
			}
			enemy.hp = 50
			continue
		}

		// Enemy turn
		fmt.Println("Enemy's turn")
		damage := enemy.attack - player.defense
		if damage < 0 {
			damage = 0
		}
		player.hp -= damage
		fmt.Println("Enemy attacks for", damage, "damage")

		// Check if player is defeated
		if player.hp <= 0 {
			fmt.Println("Player is defeated")
			break
		}
	}
}
