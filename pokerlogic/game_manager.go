package pokerlogic

import (
	"errors" // Import the errors package
)

type Player struct {
	ID    int
	Hand  []Card
	Bet   int
	Chips int
}

type GameManager struct {
	Players []Player
	Deck    *Deck
	Pot     int
}

func NewGameManager(numPlayers int) *GameManager {
	players := make([]Player, numPlayers)
	for i := range players {
		players[i] = Player{ID: i, Chips: 1000} // Default chips
	}
	return &GameManager{
		Players: players,
		Deck:    NewDeck(),
		Pot:     0,
	}
}

func (gm *GameManager) DealInitialHands(cardsPerPlayer int) error {
	for i := range gm.Players {
		hand, err := gm.Deck.Deal(cardsPerPlayer)
		if err != nil {
			return err
		}
		gm.Players[i].Hand = hand
	}
	return nil
}

func (gm *GameManager) PlaceBet(playerID, amount int) error {
	for i := range gm.Players {
		if gm.Players[i].ID == playerID {
			if gm.Players[i].Chips < amount {
				return errors.New("insufficient chips")
			}
			gm.Players[i].Bet += amount
			gm.Players[i].Chips -= amount
			gm.Pot += amount
			return nil
		}
	}
	return errors.New("player not found")
}

func (gm *GameManager) EvaluateWinner() int {
	// Placeholder logic to evaluate winner based on hands
	return 0 // Return the player ID of the winner
}
