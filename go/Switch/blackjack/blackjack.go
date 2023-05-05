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
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	}
	return 0 // "others"
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cr1, cr2, dCr := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	if cr1 == 11 && cr2 == 11 {
		return "P"
	}

	if cr1+cr2 == 21 && dCr != 11 && dCr != 0 && dCr != 10 {
		return "W"
	} else if cr1+cr2 == 21 && (dCr == 11 || dCr == 0 || dCr == 10) {
		return "S"
	}

	if v := cr1 + cr2; v >= 17 && v <= 20 {
		return "S"
	}
	if v := cr1 + cr2; v >= 12 && v <= 16 && dCr < 7 {
		return "S"
	} else if v := cr1 + cr2; v >= 12 && v <= 16 && dCr >= 7 {
		return "H"
	}

	return "H" // v <= 11
}
