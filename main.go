package main

import (
	"errors"
	"fmt"

	"github.com/neilshah12/fishtankpoker/pokerlogic" // Replace with the actual module path
)

func main() {
	// Define game modes
	// texasHoldem := pokerlogic.NewGameMode("Texas Hold'em", 2, true, false, "no limit")
	omaha := pokerlogic.NewGameMode("Omaha", 5, true, false, "pot limit")
	// badugi := pokerlogic.NewGameMode("Badugi", 4, false, true, "no limit")

	// Create a new deck
	sampleDeck := pokerlogic.NewDeck()

	// Shuffle the deck
	sampleDeck.Shuffle()

	// Deal cards for Texas Hold'em with 3 players
	playersCards, err := DealCards(sampleDeck, omaha, 9)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, cards := range playersCards {
		fmt.Printf("Player %d: %v\n", i+1, cards)
	}
}

// DealCards deals the appropriate number of cards to each player based on the game mode
func DealCards(deck *pokerlogic.Deck, gameMode *pokerlogic.GameMode, numPlayers int) ([][]pokerlogic.Card, error) {
	// generalize to 52 - num community cards instead, since it's not the case that there will be always be 5 community cards.
	if numPlayers*gameMode.CardsPerPlayer > 47 {
		return nil, errors.New("not enough cards to support this many players")
	}

	playersCards := make([][]pokerlogic.Card, numPlayers)

	for i := 0; i < numPlayers; i++ {
		cards, err := deck.Deal(gameMode.CardsPerPlayer)
		if err != nil {
			return nil, err
		}
		playersCards[i] = cards
	}

	return playersCards, nil
}
