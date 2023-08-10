package player

import "toml/blackjack/src/deck"

type Action int

var ActionStand Action = 0
var ActionHit Action = 1

type Player interface {
	DecideNextAction() Action
	GetScore() int
	GiveCard(deck.Card)
	EmptyHand()
	HasUnusedAce() bool
}
