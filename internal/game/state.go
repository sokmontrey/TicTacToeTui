package game

import (
	"TicTacToeTui/internal/cell"
	"TicTacToeTui/internal/utils/vec2"
	mapset "github.com/deckarep/golang-set/v2"
)

type State struct {
	Running     bool
	Winner      cell.Cell
	WinnerCells *mapset.Set[vec2.Vec2]
}

func NewState() State {
	return State{
		Running:     true,
		Winner:      cell.None,
		WinnerCells: nil,
	}
}

func (s *State) End(winner cell.Cell, winnerCells *mapset.Set[vec2.Vec2]) {
	s.Running = false
	s.Winner = winner
	s.WinnerCells = winnerCells
}
