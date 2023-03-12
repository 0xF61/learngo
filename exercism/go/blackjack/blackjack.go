package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "king":
		fallthrough
	case "ten":
		fallthrough
	case "queen":
		fallthrough
	case "jack":
		return 10
	default:
		return 0
	}

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myCards := ParseCard(card1) + ParseCard(card2)
	if card1 == card2 && card1 == "ace" {
		return "P"
	} else if myCards == 21 && ParseCard(dealerCard) < 10 {
		return "W"
	} else if myCards == 21 && ParseCard(dealerCard) >= 10 {
		return "S"
	} else if myCards >= 17 {
		return "S"
	} else if myCards >= 12 && ParseCard(dealerCard) >= 7 {
		return "H"
	} else if myCards >= 12 {
		return "S"
	}

	return "H"

}