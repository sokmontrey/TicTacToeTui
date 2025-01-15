package cell

import "TicTacToeTui/internal/utils/color"

type Cell byte

const (
	X Cell = iota
	O
	None
)

func (c Cell) ToString() string {
	if c == X {
		return "X"
	} else if c == O {
		return "O"
	} else {
		return "."
	}
}

func (c Cell) GetColor() string {
	if c == X {
		return color.Red
	} else if c == O {
		return color.Blue
	} else {
		return color.Gray
	}
}
