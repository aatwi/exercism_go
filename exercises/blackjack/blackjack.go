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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Number := ParseCard(card1)
	card2Number := ParseCard(card2)
	dealerNumber := ParseCard(dealerCard)
	sumOfCards := card1Number + card2Number

	if card1Number == 11 && card2Number == 11 {
		return "P"
	}
	switch {
	case sumOfCards == 21:
		if dealerNumber == 10 || dealerNumber == 11 {
			return "S"
		}
		return "W"
	case sumOfCards >= 17 && sumOfCards <= 20:
		return "S"
	case sumOfCards >= 12 && sumOfCards <= 17:
		if dealerNumber >= 7 {
			return "H"
		}
		return "S"
	case sumOfCards <= 11:
		return "H"
	}
	panic("Please implement the FirstTurn function")
}
