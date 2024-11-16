package pokerlogic

import "fmt"

func FormatCard(card Card) string {
	return fmt.Sprintf("%s of %s", card.Rank, card.Suit)
}
