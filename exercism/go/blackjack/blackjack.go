package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"king":  10,
		"ten":   10,
		"queen": 10,
		"jack":  10,
	}

	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myCardsTotal := ParseCard(card1) + ParseCard(card2)
	dealerCardTotal := ParseCard(dealerCard)

	if card1 == card2 && card1 == "ace" {
		return "P"
	} else if myCardsTotal == 21 && dealerCardTotal < 10 {
		return "W"
	} else if myCardsTotal == 21 && dealerCardTotal >= 10 {
		return "S"
	} else if myCardsTotal >= 17 {
		return "S"
	} else if myCardsTotal >= 12 && dealerCardTotal >= 7 {
		return "H"
	} else if myCardsTotal >= 12 {
		return "S"
	}

	return "H"
}
