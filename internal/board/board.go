package board

import (
	"TicTacToeTui/internal/cell"
	"TicTacToeTui/internal/utils/color"
	"TicTacToeTui/internal/utils/vec2"
	mapset "github.com/deckarep/golang-set/v2"
)

var adjDirs = [8]vec2.Vec2{
	vec2.NewVec2(-1, -1),
	vec2.NewVec2(1, 1),
	vec2.NewVec2(0, -1),
	vec2.NewVec2(0, 1),
	vec2.NewVec2(1, -1),
	vec2.NewVec2(-1, 1),
	vec2.NewVec2(-1, 0),
	vec2.NewVec2(1, 0),
}

var checkPairDirs = [4][2]vec2.Vec2{
	{vec2.NewVec2(-1, -1), vec2.NewVec2(1, 1)},
	{vec2.NewVec2(0, -1), vec2.NewVec2(0, 1)},
	{vec2.NewVec2(1, -1), vec2.NewVec2(-1, 1)},
	{vec2.NewVec2(-1, 0), vec2.NewVec2(1, 0)},
}

type Board struct {
	cells          map[vec2.Vec2]cell.Cell
	Width          int
	Height         int
	originPosition vec2.Vec2
	cursorOffset   vec2.Vec2
	isCentered     bool
}

func NewBoard(width int, height int) Board {
	return Board{
		cells:          make(map[vec2.Vec2]cell.Cell),
		Width:          width,
		Height:         height,
		originPosition: vec2.ZeroVec2(),
		cursorOffset:   vec2.ZeroVec2(),
		isCentered:     true,
	}
}

func (b *Board) ToggleCentered() {
	b.isCentered = !b.isCentered
	b.CenterBoard()
}

func (b *Board) CenterBoard() {
	if b.isCentered {
		b.originPosition = b.originPosition.Add(b.cursorOffset)
		b.cursorOffset = vec2.ZeroVec2()
	}
}

func (b *Board) MoveUp() {
	if b.isCentered || b.cursorOffset.Y <= 1-b.Height/2 {
		b.originPosition = b.originPosition.Up()
	} else {
		b.cursorOffset = b.cursorOffset.Up()
	}
}

func (b *Board) MoveDown() {
	if b.isCentered || b.cursorOffset.Y >= b.Height/2-1 {
		b.originPosition = b.originPosition.Down()
	} else {
		b.cursorOffset = b.cursorOffset.Down()
	}
}

func (b *Board) MoveLeft() {
	if b.isCentered || b.cursorOffset.X <= 1-b.Width/2 {
		b.originPosition = b.originPosition.Left()
	} else {
		b.cursorOffset = b.cursorOffset.Left()
	}
}

func (b *Board) MoveRight() {
	if b.isCentered || b.cursorOffset.X >= b.Width/2-1 {
		b.originPosition = b.originPosition.Right()
	} else {
		b.cursorOffset = b.cursorOffset.Right()
	}
}

func (b *Board) GetCell(pos vec2.Vec2) cell.Cell {
	if cellValue, ok := b.cells[pos]; ok {
		return cellValue
	}
	return cell.None
}

func (b *Board) IsAvailable() bool {
	pos := b.originPosition.Add(b.cursorOffset)
	_, taken := b.cells[pos]
	return !taken
}

func (b *Board) HasAdjacent() bool {
	pos := b.originPosition.Add(b.cursorOffset)
	if pos == vec2.ZeroVec2() {
		return true
	}
	for _, dir := range adjDirs {
		if cellValue := b.GetCell(pos.Add(dir)); cellValue != cell.None {
			return true
		}
	}
	return false
}

func (b *Board) SetCell(cellValue cell.Cell) {
	pos := b.originPosition.Add(b.cursorOffset)
	b.cells[pos] = cellValue
}

func (b *Board) MakeMove(cellPlaced cell.Cell) (bool, *mapset.Set[vec2.Vec2]) {
	b.SetCell(cellPlaced)
	placedPos := b.originPosition.Add(b.cursorOffset)
	for _, dirPair := range checkPairDirs {
		countedCells := mapset.NewSet[vec2.Vec2](placedPos)
		for _, dir := range dirPair { // only 2 runs
			tempPos := placedPos.Add(dir)
			for cellPlaced == b.GetCell(tempPos) {
				countedCells.Add(tempPos)
				tempPos = tempPos.Add(dir)
			}
		}
		if len(countedCells.ToSlice()) >= 5 {
			return true, &countedCells
		}
	}
	return false, nil
}

func (b *Board) getBracket(currentPlayer cell.Cell, offset vec2.Vec2) (string, string) {
	prefix := " "
	postfix := ""
	if offset == b.cursorOffset {
		prefix = "["
		postfix = "]"
	} else if offset.Left() == b.cursorOffset {
		prefix = ""
	}
	currentColor := currentPlayer.GetColor()
	return currentColor + prefix + color.Reset,
		currentColor + postfix + color.Reset
}

func (b *Board) getCellMark(pos vec2.Vec2, winnerCells *mapset.Set[vec2.Vec2]) string {
	c := b.GetCell(pos)
	cellIcon := ""
	if pos == vec2.ZeroVec2() && c == cell.None {
		cellIcon = "+"
	} else {
		cellIcon = c.ToString()
	}
	// Dealing with cell color
	cellColor := c.GetColor()
	if winnerCells != nil && (*winnerCells).Contains(pos) {
		cellColor = color.Yellow
	}
	return cellColor + cellIcon + color.Reset
}

func (b *Board) ToString(currentPlayer cell.Cell, winnerCells *mapset.Set[vec2.Vec2]) string {
	result := ""
	for y := 1 - b.Height/2; y < b.Height/2; y++ {
		for x := 1 - b.Width/2; x < b.Width/2; x++ {
			offset := vec2.NewVec2(x, y)
			prefix, postfix := b.getBracket(currentPlayer, offset)
			pos := b.originPosition.Add(offset)
			result += prefix + b.getCellMark(pos, winnerCells) + postfix
		}
		result += "\n"
	}
	return result
}
