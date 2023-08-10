package deck

import (
	"fmt"
	"math/rand"
)

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

type Deck struct {
	cards Stack
}

func (this *Deck) Reset() {
	// Empty deck (to save memory)
	len := this.cards.length
	for i := 0; i < len; i++ {
		this.cards.Pop()
	}

	if this.cards.length != 0 {
		fmt.Println("Warning: deck isn't empty!")
		fmt.Println(this.cards.length)
	}

	var j [52]bool

	// 52 cards in the deck havent been added yet
	for i := 0; i < 52; i++ {
		j[i] = false
	}

	// Add the 52 cards to the deck in a random order
	for i := 0; i < 52; i++ {
		chosen := false
		for chosen == false {
			r := rand.Intn(52)

			if j[r] == false {
				this.cards.Push(Card{
					revealed: true,
					index:    r,
				})
				j[r] = true
				chosen = true
			}
		}
	}
}

func (this *Deck) DrawCard() Card {
	return this.cards.Pop().(Card)
}
