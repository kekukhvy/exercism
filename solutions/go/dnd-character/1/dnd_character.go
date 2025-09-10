package dndcharacter

import (
	"log"
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
	abilitties := generateAbilityArray()
	log.Println("Generated array", abilitties)
	sort(abilitties)
	log.Println("Sorted array", abilitties)
	res := sum(abilitties)
	log.Println("Result", res)
	return res
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

func generateAbilityArray() []int {
	abilitties := make([]int, 4)

	for i := range abilitties {
		abilitties[i] = rand.Intn(6) + 1

	}

	return abilitties
}

func sort(in []int) {
	for i := 0; i < len(in)-1; i++ {
		for j := 0; j < len(in)-1-i; j++ {
			if in[j] < in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
			}
		}
	}
}

func sum(in []int) int {
	sum := 0
	for _, v := range in[:3] {
		sum += v
	}
	return sum
}

func calculateHitpoint(constitution int) int {
	return (10 + constitution) / 2
}
