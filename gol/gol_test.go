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
	field := Field{Cell{0, 0}: false, Cell{1, 1}: false}
	actual := field.neighborCounts()
	expected := Counts{Cell{2, 1}: 1, Cell{1, 0}: 2,
		Cell{1, 1}: 1, Cell{2, 0}: 1, Cell{-1, -1}: 1,
		Cell{-1, 1}: 1, Cell{0, -1}: 1, Cell{0, 1}: 2,
		Cell{1, 2}: 1, Cell{2, 2}: 1, Cell{-1, 0}: 1,
		Cell{1, -1}: 1, Cell{0, 0}: 1, Cell{0, 2}: 1}

	if len(expected) != len(actual) {
		t.Error("Size is different", expected, actual)
	}

	for cell, expectedValue := range expected {
		if actualValue, ok := actual[cell]; ok {
			if expectedValue != actualValue {
				t.Error("Values differ", expected, actual)
			}
		} else {
			t.Error("Cell not found", cell)
		}
	}
}

func TestDebugStringEmpty(t *testing.T) {
	field := Field{}
	actual := field.debugString(0)
	expected := "empty"
	if expected != actual {
		t.Error("String not as expected", expected, actual)
	}
}

func TestDebugString(t *testing.T) {
	field := Field{Cell{0, 0}: false, Cell{1, 1}: false, Cell{-1, 1}: false}
	actual := field.debugString(0)
	expected := ".X.\nX.X\n"
	if expected != actual {
		t.Error("String not as expected", expected, actual)
	}
}

func TestFromString(t *testing.T) {
	description := "X..\n.XX"
	expected := Field{Cell{0, 0}: false, Cell{1, 1}: false, Cell{2, 1}: false}
	actual := MakeField(description)
	if reflect.DeepEqual(expected, actual) {
		t.Error("Incorrect field created", expected, actual)
	}
}

func TestStepSimple(t *testing.T) {
	field := MakeField("...\nXXX\n...")
	expected := "X\nX\nX\n"
	field = field.Step()
	actual := field.debugString(0)
	if expected != actual {
		t.Error("Field advanced incorrectly", expected, actual)
	}
}

func TestStepsGlider(t *testing.T) {
	states := [...]string{".X.\n..X\nXXX\n", "X.X\n.XX\n.X.\n", "..X\nX.X\n.XX\n", "X..\n.XX\nXX.\n"}
	field := MakeField(states[0])
	for step, expected := range states {
		actual := field.debugString(0)
		if expected != actual {
			t.Error("Field advanced incorrectly in step", step, expected, actual)
		}
		field = field.Step()
	}
}

func BenchmarkGameOfLife(b *testing.B) {
	field := MakeField(".X.\n..X\nXXX")
	for n := 0; n < b.N; n++ {
		field = field.Step()
	}
}
