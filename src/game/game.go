package game

import (
	"fmt"
	"toml/blackjack/src/deck"
	"toml/blackjack/src/player"
)

type Game struct {
	Deck   deck.Deck
	AI     *player.AI
	Dealer player.Dealer
}

func (this *Game) Play() {
	// Main game loop

	fmt.Println(" =======================")
	fmt.Println(" == Starting new game ==")
	fmt.Println(" =======================")

	// Dish out cards
	this.AI.EmptyHand()
	this.Dealer.EmptyHand()
	this.Deck.Reset()

	this.AI.GiveCard(this.Deck.DrawCard())
	this.AI.GiveCard(this.Deck.DrawCard())

	dealerCard := this.Deck.DrawCard()
	this.Dealer.GiveCard(dealerCard)

	// AI needs to know the value of the dealer's card to make better decisions
	this.AI.DealersCard = dealerCard.GetValue()

	// Secret dealer card
	secretCard := this.Deck.DrawCard()
	secretCard.Hide()
	this.Dealer.GiveCard(secretCard)

	// Make a note of the AI's decision
	var decision player.Action
	var currentScore int
	var hasUnusedAce bool

	// The AI needs to make all of it's decisions
	for {
		currentScore = this.AI.GetScore()
		hasUnusedAce = this.AI.HasUnusedAce()
		decision = this.AI.DecideNextAction()

		if decision == player.ActionHit {
			// AI has decided to draw a card
			this.AI.GiveCard(this.Deck.DrawCard())

			if this.AI.GetScore() >= 22 {
				// AI lost due to drawing a card - maximum punishment
				fmt.Println("AI busts!!")
				this.AI.GiveReward(currentScore, hasUnusedAce, decision, -1)
				break
			} else {
				// AI is still in the game! Give it partial reward for not losing
				this.AI.GiveReward(currentScore, hasUnusedAce, decision, 0.25)
			}
		} else {
			// AI has decided to stop drawing cards
			break
		}
	}

	if this.DetermineWinner() != -1 && decision == player.ActionStand {
		// AI has not lost yet - we need to determine the dealer's actions now

		// Dealer reveals hidden cards
		this.Dealer.RevealCards()

		for this.Dealer.DecideNextAction() == player.ActionHit {
			this.Dealer.GiveCard(this.Deck.DrawCard())
		}

		fmt.Println("AI's final score: ", this.AI.GetScore())
		fmt.Println("Dealer's final score: ", this.Dealer.GetScore())

		// No more cards can be dealt - let's see who the winner is!
		if this.DetermineWinner() == 1 {
			// Player has won - reward them for standing at the correct moment
			fmt.Println("AI wins!!")
			this.AI.GiveReward(currentScore, hasUnusedAce, decision, 1)
		} else if this.DetermineWinner() == -1 {
			// Player has lost - punish them for standing
			fmt.Println("Dealer wins!!")
			this.AI.GiveReward(currentScore, hasUnusedAce, decision, -1)
		} else {
			// Draw - do not give reward
			fmt.Println("Draw")
		}
	}

	// End of game
	fmt.Println("End of game cleanup")
	this.AI.Episode()
}

func (this *Game) DetermineWinner() int {
	if this.AI.GetScore() >= 22 {
		return -1 // AI busts
	} else if this.Dealer.GetScore() >= 22 {
		return 1 // Dealer busts
	} else if this.AI.GetScore() > this.Dealer.GetScore() {
		return 1 // AI has better score
	} else if this.AI.GetScore() < this.Dealer.GetScore() {
		return -1 // Dealer has better score
	} else {
		return 0 // Neither dealer or player has bust, and both have same score
	}
}
