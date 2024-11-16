package main

import (
	"fmt"

	"github.com/neilshah12/fishtankpoker/pokerlogic" // Replace with the actual module path
)

func main() {
	// Create a new deck
	sampleDeck := pokerlogic.NewDeck()

	// Print the deck
	for _, card := range sampleDeck.Cards {
		fmt.Println(pokerlogic.FormatCard(card))
	}
}
