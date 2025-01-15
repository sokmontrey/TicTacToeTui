package game

import (
	"TicTacToeTui/internal/board"
	"TicTacToeTui/internal/cell"
	"TicTacToeTui/internal/utils/color"
	tea "github.com/charmbracelet/bubbletea"
)

type Game struct {
	currentPlayer cell.Cell
	board         board.Board
	msg           Msg
	state         State
}

func NewGame() Game {
	return Game{
		currentPlayer: cell.X,
		board:         board.NewBoard(20, 10),
		msg:           NewEmptyMsg(),
		state:         NewState(),
	}
}

func (g Game) Restart() {
	width := g.board.Width
	height := g.board.Height
	g.currentPlayer = cell.X
	g.board = board.NewBoard(width, height)
	g.msg = NewEmptyMsg()
	g.state = NewState()
}

func (g Game) Init() tea.Cmd {
	return nil
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	g.msg = NewEmptyMsg()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return g, tea.Quit
		case "w", "up":
			g.board.MoveUp()
		case "s", "down":
			g.board.MoveDown()
		case "a", "left":
			g.board.MoveLeft()
		case "d", "right":
			g.board.MoveRight()
		case "c":
			g.board.ToggleCentered()
		case "r":
			g.Restart()
		case " ", "enter":
			if !g.state.Running {
				return g, nil
			} else if !g.board.IsAvailable() {
				g.msg = NewErrorMsg("Cell taken")
			} else if !g.board.HasAdjacent() {
				g.msg = NewErrorMsg("Too far apart!")
			} else if won, winnerCells := g.board.MakeMove(g.currentPlayer); won {
				g.state.End(g.currentPlayer, winnerCells)
			} else {
				g.currentPlayer = swapPlayer(g.currentPlayer)
			}
		}
	}
	return g, nil
}

func (g Game) View() string {
	s := "\n"
	s += color.Gray + "   W           C: toggle centered\n" + color.Reset
	s += color.Gray + " A S D: move   R: restart game\n\n " + color.Reset

	if !g.state.Running {
		s += color.Green + g.state.Winner.ToString() + " won!" + color.Reset
	} else if g.msg.IsEmpty {
		s += g.currentPlayer.GetColor() + g.currentPlayer.ToString() + " to move"
	} else if g.msg.IsError {
		s += color.Red + g.msg.Value + color.Reset
	} else {
		s += color.Green + g.msg.Value + color.Reset
	}

	s += "\n\n"
	s += g.board.ToString(g.currentPlayer, g.state.WinnerCells)
	return s
}

func swapPlayer(currentPlayer cell.Cell) cell.Cell {
	if currentPlayer == cell.X {
		return cell.O
	} else {
		return cell.X
	}
}
