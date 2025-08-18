package dndcharacter

import (
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	diff := score - 10
	if diff >= 0 {
		return diff / 2
	}

	if diff%2 == 0 {
		return diff / 2
	}
	return diff/2 - 1
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return rand.Intn(15) + 3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	char := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	char.Hitpoints = calculateHitpoint(char.Constitution)
	return char
}

func calculateHitpoint(constitution int) int {
	return (10 + constitution) / 2
}
