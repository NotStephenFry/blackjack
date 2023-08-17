package player

import (
	"fmt"
	"math"
	"math/rand"
	"toml/blackjack/src/deck"
)

type AI struct {
	cards          []*deck.Card
	qTable         [25][11][2][2]float64
	DealersCard    int
	LearningRate   float64
	DecayRate      float64
	DiscountFactor float64
	ExploitRate    float64
	Epsilon        float64
}

func (this *AI) Initialise() {
	// Set all possible actions to 0 expected reward

	// Possible player card value
	for i := 1; i <= 23; i++ {

		// Revealed dealer card
		for j := 0; j <= 10; j++ {

			// Has usable ace
			for k := 0; k <= 1; k++ {

				// Possible actions
				for l := 0; l <= 1; l++ {
					this.qTable[i][j][k][l] = 0
				}
			}
		}
	}
}

func (this *AI) DecideNextAction() Action {
	d := rand.Float64()
	fmt.Println("Dec is ", d)

	fmt.Println("== AI is deciding what to do ==")
	fmt.Println("AI currently has score of ", this.GetScore())
	// Determine whether we should exploit or explore
	fmt.Println("Epsilon is ", this.Epsilon)

	exp := rand.Float64()
	fmt.Println("Rand is ", exp)

	if this.Epsilon > exp {
		fmt.Println("AI chose to exploit")
		// Exploitation - choose best action
		/*ace := 0
		if this.HasUnusedAce() {
			ace = 1
		}*/

		rewardHit := this.qTable[this.GetScore()][this.DealersCard][0][ActionHit]
		fmt.Println("AI thinks hitting will give ", rewardHit, " reward")
		rewardStand := this.qTable[this.GetScore()][this.DealersCard][0][ActionStand]
		fmt.Println("AI thinks standing will give ", rewardStand, " reward")

		if rewardHit > rewardStand {
			fmt.Println("AI chose to hit")
			return ActionHit
		} else if rewardHit < rewardStand {
			fmt.Println("AI chose to stand")
			return ActionStand
		}

		fmt.Println("AI was indecisive exploiting")
	}

	// Exploration - random action
	fmt.Println("AI chose to explore")

	if d > float64(0.5) {
		fmt.Println("AI chose to stand")
		return ActionStand
	}

	fmt.Println("AI chose to hit")
	return ActionHit
}

func (this *AI) GetScore() int {
	score := 0
	aces := 0

	for _, card := range this.cards {
		val := card.GetValue()

		if val == 1 {
			aces++
		} else {
			score += val
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

func (this *AI) GiveCard(c deck.Card) {
	fmt.Println("AI draws card: ", c.GetValue())
	this.cards = append(this.cards, &c)
}

func (this *AI) EmptyHand() {
	this.cards = nil
}

func (this *AI) HasUnusedAce() bool {
	usableAce := false
	aces := 0
	score := 0

	// Determine how many aces we have
	// And tally up our score WITHOUT aces
	for _, card := range this.cards {
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

func (this *AI) lookup(playerScore int, dealerScore int, hasUsableAce int, decision int) float64 {
	p := int(math.Min(22, float64(playerScore)))

	// TODO: in calculating the maximum expected reward, we need to predict a maximum punishment if the dealt card would bust the AI
	if p == 22 {
		return -20
	}

	return this.qTable[p][dealerScore][0][decision]
}

func (this *AI) calculateExpectedHitReward(playerScore int, dealerScore int, ace int) float64 {

	// The problem is that after taking a card, it's not guaranteed which state we will be in after taking an action
	// The solution I have devised is to take the mean average of the expected rewards of all possible future states
	// J Q K count at 10, so 10 is counted a total of 4 times in the mean average

	// All possible future rewards after hitting in the next state
	rewardIf2Take := this.lookup(playerScore+2, dealerScore, ace, 1)
	rewardIf3Take := this.lookup(playerScore+3, dealerScore, ace, 1)
	rewardIf4Take := this.lookup(playerScore+4, dealerScore, ace, 1)
	rewardIf5Take := this.lookup(playerScore+5, dealerScore, ace, 1)
	rewardIf6Take := this.lookup(playerScore+6, dealerScore, ace, 1)
	rewardIf7Take := this.lookup(playerScore+7, dealerScore, ace, 1)
	rewardIf8Take := this.lookup(playerScore+8, dealerScore, ace, 1)
	rewardIf9Take := this.lookup(playerScore+9, dealerScore, ace, 1)
	rewardIf10Take := this.lookup(playerScore+10, dealerScore, ace, 1)

	aceScore := playerScore
	if aceScore <= 10 {
		aceScore += 11
	} else {
		aceScore += 1
	}

	rewardIfAceTake := this.lookup(aceScore, dealerScore, ace, 1)

	// All possible future rewards after standing in the next state
	rewardIf2Stand := this.lookup(playerScore+2, dealerScore, ace, 0)
	rewardIf3Stand := this.lookup(playerScore+3, dealerScore, ace, 0)
	rewardIf4Stand := this.lookup(playerScore+4, dealerScore, ace, 0)
	rewardIf5Stand := this.lookup(playerScore+5, dealerScore, ace, 0)
	rewardIf6Stand := this.lookup(playerScore+6, dealerScore, ace, 0)
	rewardIf7Stand := this.lookup(playerScore+7, dealerScore, ace, 0)
	rewardIf8Stand := this.lookup(playerScore+8, dealerScore, ace, 0)
	rewardIf9Stand := this.lookup(playerScore+9, dealerScore, ace, 0)
	rewardIf10Stand := this.lookup(playerScore+10, dealerScore, ace, 0)
	rewardIfAceStand := this.lookup(aceScore, dealerScore, ace, 0)

	// For each possible card that could be drawn, determine the best action in that state
	rewardIf2 := math.Max(rewardIf2Stand, rewardIf2Take)
	rewardIf3 := math.Max(rewardIf3Stand, rewardIf3Take)
	rewardIf4 := math.Max(rewardIf4Stand, rewardIf4Take)
	rewardIf5 := math.Max(rewardIf5Stand, rewardIf5Take)
	rewardIf6 := math.Max(rewardIf6Stand, rewardIf6Take)
	rewardIf7 := math.Max(rewardIf7Stand, rewardIf7Take)
	rewardIf8 := math.Max(rewardIf8Stand, rewardIf8Take)
	rewardIf9 := math.Max(rewardIf9Stand, rewardIf9Take)
	rewardIf10 := math.Max(rewardIf10Stand, rewardIf10Take)
	rewardIfAce := math.Max(rewardIfAceStand, rewardIfAceTake)

	// Calculate the mean expected reward in future states
	var averageReward float64 = 0
	averageReward += rewardIf2 / 13
	averageReward += rewardIf3 / 13
	averageReward += rewardIf4 / 13
	averageReward += rewardIf5 / 13
	averageReward += rewardIf6 / 13
	averageReward += rewardIf7 / 13
	averageReward += rewardIf8 / 13
	averageReward += rewardIf9 / 13
	averageReward += (4 * rewardIf10 / 13)
	averageReward += rewardIfAce / 13

	return averageReward
}

func (this *AI) GiveReward(oldScore int, hadUsableAce bool, decision Action, reward float64) {
	/*ace := 0
	if hadUsableAce == true {
		ace = 1
	}*/

	currentReward := this.qTable[oldScore][this.DealersCard][0][decision]
	var nextReward float64

	if decision == ActionStand {
		// Expected reward by staying in the current state
		nextReward = currentReward
	} else {
		nextReward = this.calculateExpectedHitReward(oldScore, this.DealersCard, 0)
	}

	fmt.Println("AI had score of: ", oldScore, ". Dealer's card was: ", this.DealersCard, ". Had unused ace: ", 0, ". Decision was to: ", decision)
	fmt.Println("Current reward for this state/action: ", currentReward)
	fmt.Println("Expected future reward: ", nextReward)
	fmt.Println("Actual reward for this state/action: ", reward)

	// Bellman equation to update the qtable
	this.qTable[oldScore][this.DealersCard][0][decision] = ((1 - this.LearningRate) * currentReward) + (this.LearningRate * (reward + (this.DiscountFactor * nextReward)))

	fmt.Println("Updated reward for this state/action: ", this.qTable[oldScore][this.DealersCard][0][decision])

	// Decay the learning rate
	this.LearningRate *= this.DecayRate
}

func (this *AI) Episode() {
	fmt.Println(this.LearningRate)
	this.Epsilon += this.ExploitRate
}

func (this *AI) DumpBrain() {
	for i := 1; i <= 22; i++ {

		// Revealed dealer card
		for j := 0; j <= 10; j++ {

			// Has usable ace
			outcome := "???"

			rewardStand := this.qTable[i][j][0][ActionStand]
			rewardHit := this.qTable[i][j][0][ActionHit]

			if rewardStand > rewardHit {
				outcome = "STAND"
			} else if rewardStand < rewardHit {
				outcome = "HIT"
			}

			fmt.Println("AI had score of: ", i, ". Dealer's card was: ", j, ". Best action is ", outcome, " (", rewardStand, " VS ", rewardHit, ")")
		}
	}
}
