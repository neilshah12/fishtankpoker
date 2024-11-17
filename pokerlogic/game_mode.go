package pokerlogic

type GameMode struct {
	Name             string
	CardsPerPlayer   int
	CommunityCards   bool
	Drawing          bool
	BettingStructure string // "pot limit" or "no limit"
}

func NewGameMode(name string, cardsPerPlayer int, communityCards, drawing bool, bettingStructure string) *GameMode {
	return &GameMode{
		Name:             name,
		CardsPerPlayer:   cardsPerPlayer,
		CommunityCards:   communityCards,
		Drawing:          drawing,
		BettingStructure: bettingStructure,
	}
}
