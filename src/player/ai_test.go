package player

import (
	"testing"
	"toml/blackjack/src/deck"
)

func Test_AICountsCorrectly(t *testing.T) {
	d := deck.Deck{}
	d.Reset()

	firstCard := d.DrawCard()
	firstCard.SetIndex(6) // this is a 3

	secondCard := d.DrawCard()
	secondCard.SetIndex(45) // this is a 10

	ai := AI{}
	ai.GiveCard(firstCard)
	ai.GiveCard(secondCard)

	if ai.GetScore() != 13 {
		t.Errorf("AI does not count hand card correctly")
	}
}

func Test_AICountsAcesCorrectly(t *testing.T) {
	d := deck.Deck{}
	d.Reset()

	firstCard := d.DrawCard()
	firstCard.SetIndex(6) // this is a 3

	secondCard := d.DrawCard()
	secondCard.SetIndex(50) // this is an ace

	ai := AI{}
	ai.GiveCard(firstCard)
	ai.GiveCard(secondCard)

	// If an ace can be high OR low, it should be treated high
	// 3 + 11
	if ai.GetScore() != 14 {
		t.Errorf("AI does not count aces correctly, expected %d, got %d", 14, ai.GetScore())
	}

	thirdCard := d.DrawCard()
	thirdCard.SetIndex(45) // this is a 10
	ai.GiveCard(thirdCard)

	// If a high ace scored greater than 21, it should be treated low
	// 3 + 1 + 10
	if ai.GetScore() != 14 {
		t.Errorf("AI does not count aces correctly, expected %d, got %d", 14, ai.GetScore())
	}

	fourthCard := d.DrawCard()
	fourthCard.SetIndex(51) // this is another ace
	ai.GiveCard(fourthCard)

	// If a high ace scored greater than 21, it should be treated low
	// 3 + 1 + 10 + 1
	if ai.GetScore() != 15 {
		t.Errorf("AI does not count aces correctly, expected %d, got %d", 15, ai.GetScore())
	}
}

func Test_DealerCountsCorrectly(t *testing.T) {
	d := deck.Deck{}
	d.Reset()

	firstCard := d.DrawCard()
	firstCard.SetIndex(6) // this is a 3

	secondCard := d.DrawCard()
	secondCard.SetIndex(45) // this is a 10

	dealer := Dealer{}
	dealer.GiveCard(firstCard)
	dealer.GiveCard(secondCard)

	if dealer.GetScore() != 13 {
		t.Errorf("Dealer does not count hand cards correctly")
	}
}

func Test_DealerCountsAcesCorrectly(t *testing.T) {
	d := deck.Deck{}
	d.Reset()

	firstCard := d.DrawCard()
	firstCard.SetIndex(11) // this is a 4

	secondCard := d.DrawCard()
	secondCard.SetIndex(50) // this is an ace

	dealer := Dealer{}
	dealer.GiveCard(firstCard)
	dealer.GiveCard(secondCard)

	// If an ace can be high OR low, it should be treated high
	// 4 + 11
	if dealer.GetScore() != 15 {
		t.Errorf("AI does not count aces correctly, expected %d, got %d", 15, dealer.GetScore())
	}

	thirdCard := d.DrawCard()
	thirdCard.SetIndex(27) // this is a 8
	dealer.GiveCard(thirdCard)

	// If a high ace scored greater than 21, it should be treated low
	// 4 + 1 + 8
	if dealer.GetScore() != 13 {
		t.Errorf("AI does not count aces correctly, expected %d, got %d", 13, dealer.GetScore())
	}

	fourthCard := d.DrawCard()
	fourthCard.SetIndex(51) // this is another ace
	dealer.GiveCard(fourthCard)

	// If a high ace scored greater than 21, it should be treated low
	// 4 + 1 + 8 + 1
	if dealer.GetScore() != 14 {
		t.Errorf("AI does not count aces correctly, expected %d, got %d", 14, dealer.GetScore())
	}
}

func Test_DealerCountsHiddenCardsCorrectly(t *testing.T) {
	d := deck.Deck{}
	d.Reset()

	firstCard := d.DrawCard()
	firstCard.SetIndex(11) // this is a 4

	hiddenCard := d.DrawCard()
	hiddenCard.SetIndex(27) // this is a 8
	hiddenCard.Hide()

	dealer := Dealer{}
	dealer.GiveCard(firstCard)
	dealer.GiveCard(hiddenCard)

	// 4 + hidden
	if dealer.GetScore() != 4 {
		t.Errorf("AI does not count hidden cards correctly, expected %d, got %d", 4, dealer.GetScore())
	}

	// and now we reveal the existing cards!
	dealer.RevealCards()

	// 4 + 8
	if dealer.GetScore() != 12 {
		t.Errorf("AI does not count hidden cards correctly, expected %d, got %d", 12, dealer.GetScore())
	}
}
