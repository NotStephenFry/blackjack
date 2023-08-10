package player

import (
	"fmt"
	"toml/blackjack/src/deck"
)

type Dealer struct {
	cards []*deck.Card
}

// Dealer must stand on 17 and must draw to 16
func (this *Dealer) DecideNextAction() Action {
	fmt.Println("Dealer's score is ", this.GetScore())
	if this.GetScore() >= 17 {
		return ActionStand
	}

	return ActionHit
}

func (this *Dealer) RevealCards() {
	for _, card := range this.cards {
		card.Reveal()
	}
}

func (this *Dealer) GetScore() int {
	score := 0
	aces := 0

	for _, card := range this.cards {
		if card.IsRevealed() {

			val := card.GetValue()

			if val == 1 {
				aces++
			} else {
				score += val
			}
		}
	}

	for i := 0; i < aces; i++ {
		if (score + 11) > 21 {
			// High ace busts us - make it low instead
			score += 1
		} else {
			// High ace is fine
			score += 11
		}

	}

	return score
}

func (this *Dealer) GiveCard(c deck.Card) {
	fmt.Println("Dealer draws card: ", c.GetValue())
	this.cards = append(this.cards, &c)
}

func (this *Dealer) EmptyHand() {
	this.cards = nil
}

func (this *Dealer) HasUnusedAce() bool {
	usableAce := false
	aces := 0
	score := 0

	// Determine how many aces we have has
	// And tally up our score WITHOUT aces
	for _, card := range this.cards {
		if !card.IsRevealed() {
			continue
		}

		if card.GetValue() == 1 {
			aces++
		} else {
			score += card.GetValue()
		}
	}

	// Now see how many of our aces we can score as 11
	for i := 0; i < aces; i++ {
		if score+11 > 21 {
			// High ace will bust us - treat it as low ace instead
			score++
		} else {
			// High ace will NOT bust us!!
			usableAce = true
			break
		}
	}

	return usableAce
}
