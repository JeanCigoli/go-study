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
	cardValueOne := ParseCard(card1)
	cardValueTwo := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	sumCards := cardValueOne + cardValueTwo

	switch {
		case cardValueOne == 11 && cardValueTwo == 11:
			return "P";
		case sumCards == 21 && dealerCardValue < 10:
			return "W"
		case sumCards == 21 && dealerCardValue >= 10:
			return "S"
		case sumCards >= 17 && sumCards <= 20:
			return "S"
		case sumCards >= 12 && sumCards <= 26 && dealerCardValue >= 7:
			return "H"
		case sumCards >= 12 && sumCards <= 26 && dealerCardValue < 7:
			return "S"
		case sumCards <= 11:
			return "H"
		default:
			return ""
	}
}
