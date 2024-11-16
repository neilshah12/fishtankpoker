package pokerlogic

import "errors"

var (
	ErrNotEnoughCards = errors.New("not enough cards in the deck")
	ErrInvalidBet     = errors.New("invalid bet amount")
)
