package pokerlogic

import (
	"errors"
	"math/rand"
	"time"
)

var (
	Suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	Ranks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
)

type Card struct {
	Rank string
	Suit string
}

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	var cards []Card
	for _, suit := range Suits {
		for _, rank := range Ranks {
			cards = append(cards, Card{Rank: rank, Suit: suit})
		}
	}
	return &Deck{Cards: cards}
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func (d *Deck) Deal(numCards int) ([]Card, error) {
	if len(d.Cards) < numCards {
		return nil, errors.New("not enough cards in the deck")
	}
	dealt := d.Cards[:numCards]
	d.Cards = d.Cards[numCards:]
	return dealt, nil
}

