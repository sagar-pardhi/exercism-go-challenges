package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
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
	// panic("Please implement the FirstTurn function")
	// stand = (S)
	// Hit = (H)
	// Split = (P)
	// Win = (W)

	switch {
	// if pair of ace
	case card1 == "ace" && card2 == "ace":
		return "P"
	// if blackjack
	case (ParseCard(card1)+ParseCard(card2) == 21) && (dealerCard != "ace") && (ParseCard(dealerCard) != 10):
		return "W"
	case (ParseCard(card1)+ParseCard(card2) == 21) && (ParseCard(dealerCard) >= 10) && (ParseCard(dealerCard) <= 11):
		return "S"
	// if sum of cards in range [17, 20]
	case (ParseCard(card1)+ParseCard(card2) >= 17) && (ParseCard(card1)+ParseCard(card1) <= 20):
		return "S"
	// if sum of cards in range [12, 16]
	case (ParseCard(card1)+ParseCard(card2) >= 12) && (ParseCard(card1)+ParseCard(card2) <= 16) && (ParseCard(dealerCard) < 7):
		return "S"
	case (ParseCard(card1)+ParseCard(card2) >= 12) && (ParseCard(card1)+ParseCard(card1) <= 16) && (ParseCard(dealerCard) >= 7):
		return "H"
	// if sum of cards <= 11
	case ParseCard(card1)+ParseCard(card2) <= 11:
		return "H"
	default:
		return "H"
	}
}
