package deck

type Card struct {
	revealed bool
	index    int
}

func (this *Card) IsRevealed() bool {
	return this.revealed
}

func (this *Card) score() int {
	switch this.index {
	// 2s
	case 0:
		return 2
	case 1:
		return 2
	case 2:
		return 2
	case 3:
		return 2

	// 3s
	case 4:
		return 3
	case 5:
		return 3
	case 6:
		return 3
	case 7:
		return 3

	// 4s
	case 8:
		return 4
	case 9:
		return 4
	case 10:
		return 4
	case 11:
		return 4

	// 5s
	case 12:
		return 5
	case 13:
		return 5
	case 14:
		return 5
	case 15:
		return 5

	// 6s
	case 16:
		return 6
	case 17:
		return 6
	case 18:
		return 6
	case 19:
		return 6

	// 7s
	case 20:
		return 7
	case 21:
		return 7
	case 22:
		return 7
	case 23:
		return 7

	// 8s
	case 24:
		return 8
	case 25:
		return 8
	case 26:
		return 8
	case 27:
		return 8

	// 9s
	case 28:
		return 9
	case 29:
		return 9
	case 30:
		return 9
	case 31:
		return 9

	// 10s
	case 32:
		return 10
	case 33:
		return 10
	case 34:
		return 10
	case 35:
		return 10

	// Js
	case 36:
		return 10
	case 37:
		return 10
	case 38:
		return 10
	case 39:
		return 10

	// Qs
	case 40:
		return 10
	case 41:
		return 10
	case 42:
		return 10
	case 43:
		return 10

	// Ks
	case 44:
		return 10
	case 45:
		return 10
	case 46:
		return 10
	case 47:
		return 10

	// Special: Aces
	case 48:
		return 1
	case 49:
		return 1
	case 50:
		return 1
	case 51:
		return 1

	// should never happen
	default:
		return -1
	}
}

func (this *Card) GetValue() int {
	if this.IsRevealed() {
		return this.score()
	}

	return 0
}

func (this *Card) Reveal() {
	this.revealed = true
}

func (this *Card) Hide() {
	this.revealed = false
}

// Used only in the tests
func (this *Card) GetIndex() int {
	return this.index
}

func (this *Card) SetIndex(i int) {
	this.index = i
}
