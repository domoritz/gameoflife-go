package gol

import (
	"reflect"
	"testing"
)

func TestMakeCell(t *testing.T) {
	x := 1
	y := 2
	c := Cell{x, y}
	if c.X != x {
		t.Error("x is not as expected", x, c.X)
	}
	if c.Y != y {
		t.Error("y is not as expected", y, c.Y)
	}
}

func TestNeighbors(t *testing.T) {
	cell := Cell{0, 1}
	var actual [8]Cell
	expected := [8]Cell{{-1, 0}, {-1, 1}, {-1, 2},
		{0, 0}, {0, 2},
		{1, 0}, {1, 1}, {1, 2}}

	cell.neighbors(&actual)
	if expected != actual {
		t.Error("Neighbors are not as expected", expected, actual)
	}
}

func TestNeighborCounts(t *testing.T) {
	board := Board{Cell{0, 0}: false, Cell{1, 1}: false}
	actual := board.neighborCounts()
	expected := Counts{Cell{2, 1}: 1, Cell{1, 0}: 2,
		Cell{1, 1}: 1, Cell{2, 0}: 1, Cell{-1, -1}: 1,
		Cell{-1, 1}: 1, Cell{0, -1}: 1, Cell{0, 1}: 2,
		Cell{1, 2}: 1, Cell{2, 2}: 1, Cell{-1, 0}: 1,
		Cell{1, -1}: 1, Cell{0, 0}: 1, Cell{0, 2}: 1}

	if len(expected) != len(actual) {
		t.Error("Size is different", expected, actual)
	}

	for cell, expected_value := range expected {
		if actual_value, ok := actual[cell]; ok {
			if expected_value != actual_value {
				t.Error("Values differ", expected, actual)
			}
		} else {
			t.Error("Cell not found", cell)
		}
	}
}

func TestDebugStringEmpty(t *testing.T) {
	board := Board{}
	actual := board.DebugString()
	expected := "empty"
	if expected != actual {
		t.Error("String not as expected", expected, actual)
	}
}

func TestDebugString(t *testing.T) {
	board := Board{Cell{0, 0}: false, Cell{1, 1}: false, Cell{-1, 1}: false}
	actual := board.DebugString()
	expected := ".X.\nX.X\n"
	if expected != actual {
		t.Error("String not as expected", expected, actual)
	}
}

func TestFromString(t *testing.T) {
	description := "X..\n.XX"
	expected := Board{Cell{0, 0}: false, Cell{1, 1}: false, Cell{2, 1}: false}
	actual := MakeBoard(description)
	if reflect.DeepEqual(expected, actual) {
		t.Error("Incorrect board created", expected, actual)
	}
}

func TestAdvanceSimple(t *testing.T) {
	board := MakeBoard("...\nXXX\n...")
	expected := "X\nX\nX\n"
	board.Advance()
	actual := board.DebugString()
	if expected != actual {
		t.Error("Board advanced incorrectly", expected, actual)
	}
}

func TestAdvanceGlider(t *testing.T) {
	states := [...]string{".X.\n..X\nXXX\n", "X.X\n.XX\n.X.\n", "..X\nX.X\n.XX\n", "X..\n.XX\nXX.\n"}
	board := MakeBoard(states[0])
	for step, expected := range states {
		actual := board.DebugString()
		if expected != actual {
			t.Error("Board advanced incorrectly in step", step, expected, actual)
		}
		board.Advance()
	}
}

func BenchmarkGameOfLife(b *testing.B) {
	board := MakeBoard(".X.\n..X\nXXX")
	for n := 0; n < b.N; n++ {
		board.Advance()
	}
}
