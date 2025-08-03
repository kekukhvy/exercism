package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	mapCardValues := map[string]int{
		"ace": 11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6,
		"seven": 7, "eight": 8, "nine": 9, "ten": 10, "jack": 10, "queen": 10,
		"king": 10, "other": 0,
	}

	return mapCardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	playerValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	switch {
	case playerValue == 21 && dealerValue >= 10:
		return "S"
	case playerValue == 21:
		return "W"
	case playerValue >= 17 && playerValue < 21:
		return "S"
	case playerValue >= 12 && playerValue <= 16 && dealerValue < 7:
		return "S"
	case playerValue >= 12 && playerValue <= 16:
		return "H"

	case playerValue < 12:
		return "H"
	}

	return "H" // Default case, should not be reached
}
