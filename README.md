# blackjack
Blackjack AI implemented in Go using Q-Learning

# What is this
In the game of Blackjack, the goal is to get as close to a score of 21 as possible. Going OVER 21 causes the player to immediately lose (Bust)

Each player is dealt 2 cards (including the dealer). One of the dealer's cards at the start is hidden.

An Ace can count as either "1" or "11"

The AI will play against a robot dealer, and the AI will make it's decisions first.

If the AI busts, the game is lost without revealing what card(s) the dealer has.

Once the AI has made it's decisions (if it hasn't lost already), the dealer will reveal it's card and then follow the same pattern every game - keep drawing until it has a score of AT LEAST 17

Once it has been determined no more cards can be drawn, the game outcome is decided by either:

- Dealer busts: **AI WINS**
- Dealer has a GREATER score than the AI: **DEALER WINS**
- Dealer has a SMALLER score than the AI: **AI WINS**
- Dealer has SAME score as the AI: **DRAW**

# How it works
The AI will take note of the current game state (It's score, dealer's score, and whether or not it has a soft Ace available)

In any given state, the AI can either Hit (draw a card), or Stand (stop drawing cards).

It will either get punished or rewarded, depending on whether a given action in a given state makes it lose the game or win the game respectively.

It will store these rewards in an internal Q-Table.

Each time it gets rewarded, it will update the value in the table based on a combination of the old reward value for the current state-action, the expected future reward for the next state(s), and the actual reward given.

At first, the AI will perform random actions, the "explore phase". But over time, it will start to build a model of which actions are VERY GOOD, and which actions are VERY BAD

Eventually, the AI will switch to the "exploit phase", and use the action with the highest Q-Score for a given state to give it what IT thinks is the best possible action in a current state.

After several tens/hundreds of thousands of games against the robot dealer, it *should* start making good decisions on how to play the game




Also bundled are several unit tests that show that the seperate parts of the game are implemented correctly

# How to run
`go run src/main.go` (optionally pipe to a txt file)
