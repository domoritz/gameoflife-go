package gol

import (
	"bytes"
	"math"
	"strings"
)

type Cell struct {
	X int
	Y int
}

type Board map[Cell]bool
type Counts map[Cell]int

// Writes the neighbors of a given cell into neighbors.
func (cell *Cell) neighbors(neighbors *[8]Cell) {
	i := 0
	for x := cell.X - 1; x < cell.X+2; x++ {
		for y := cell.Y - 1; y < cell.Y+2; y++ {
			if x != cell.X || y != cell.Y {
				neighbors[i] = Cell{x, y}
				i++
			}
		}
	}
}

// Returns a map with the neighbor counts
func (board *Board) neighborCounts() Counts {
	counts := make(Counts)
	var neighbors [8]Cell
	for cell, _ := range *board {
		cell.neighbors(&neighbors)
		for _, neighbor := range neighbors {
			if _, ok := counts[neighbor]; ok {
				counts[neighbor] += 1
			} else {
				counts[neighbor] = 1
			}
		}
	}
	return counts
}

func (board *Board) isAlive(cell Cell) bool {
	_, found := (*board)[cell]
	return found
}

// Advances the board by one step
func (board *Board) Advance() {
	new_board := make(Board)
	for cell, count := range board.neighborCounts() {
		if count == 3 || board.isAlive(cell) && count == 2 {
			new_board[cell] = true
		}
	}
	*board = new_board
}

// Creates a board from a string description.
func MakeBoard(description string) Board {
	board := make(Board)
	for y, line := range strings.Split(description, "\n") {
		for x, c := range line {
			if c == 'X' {
				board[Cell{x, y}] = true
			}
		}
	}
	return board
}

// Returns a string representation of the infinite without padding around
// cells that are alive board or "empty" if the board is empty.
func (board *Board) DebugString() string {
	if len(*board) == 0 {
		return "empty"
	}

	var buffer bytes.Buffer
	minx := math.MaxInt32
	maxx := math.MinInt32
	miny := math.MaxInt32
	maxy := math.MinInt32
	for cell, _ := range *board {
		if cell.X < minx {
			minx = cell.X
		}
		if cell.X > maxx {
			maxx = cell.X
		}
		if cell.Y < miny {
			miny = cell.Y
		}
		if cell.Y > maxy {
			maxy = cell.Y
		}
	}

	for y := miny; y < maxy+1; y++ {
		for x := minx; x < maxx+1; x++ {
			if _, ok := (*board)[Cell{x, y}]; ok {
				buffer.WriteString("X")
			} else {
				buffer.WriteString(".")
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
