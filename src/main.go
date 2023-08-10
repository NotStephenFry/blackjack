package main

import (
	"fmt"
	"toml/blackjack/src/deck"
	"toml/blackjack/src/game"
	"toml/blackjack/src/player"
)

func main() {
	deck := deck.Deck{}
	deck.Reset()

	ai := player.AI{
		LearningRate:   0.1,
		DecayRate:      0.999999,
		DiscountFactor: 0.75,
		Epsilon:        0,
		ExploitRate:    0.00001,
	}
	ai.Initialise()

	dealer := player.Dealer{}

	game := game.Game{
		Deck:   deck,
		AI:     ai,
		Dealer: dealer,
	}

	fmt.Println("starting main game loop")
	gameNumber := 0
	for {
		gameNumber++
		fmt.Println("Game Number: ", gameNumber)
		game.Play()
	}
}
