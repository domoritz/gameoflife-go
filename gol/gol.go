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

type Field map[Cell]bool
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
func (field *Field) neighborCounts() Counts {
	counts := make(Counts)
	var neighbors [8]Cell
	for cell, _ := range *field {
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

// Returns true if the cell is alive.
func (field *Field) isAlive(cell Cell) bool {
	_, found := (*field)[cell]
	return found
}

// Advances the field by one step
func (field *Field) Step() {
	new_field := make(Field)
	for cell, count := range field.neighborCounts() {
		if count == 3 || field.isAlive(cell) && count == 2 {
			new_field[cell] = true
		}
	}
	*field = new_field
}

// Creates a field from a string description.
func MakeField(description string) Field {
	field := make(Field)
	for y, line := range strings.Split(description, "\n") {
		for x, c := range line {
			if c == 'X' {
				field[Cell{x, y}] = true
			}
		}
	}
	return field
}

// Returns a string representation of the infinite with the provided padding
// around cells that are alive field or "empty" if the field is empty.
func (field *Field) DebugString(padding int) string {
	if len(*field) == 0 {
		return "empty"
	}

	var buffer bytes.Buffer
	minx := math.MaxInt32
	maxx := math.MinInt32
	miny := math.MaxInt32
	maxy := math.MinInt32
	for cell, _ := range *field {
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

	for y := miny - padding; y < maxy+1+padding; y++ {
		for x := minx - padding; x < maxx+1+padding; x++ {
			if _, ok := (*field)[Cell{x, y}]; ok {
				buffer.WriteString("X")
			} else {
				buffer.WriteString(".")
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}