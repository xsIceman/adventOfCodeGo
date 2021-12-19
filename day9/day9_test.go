package day9_test

import (
	"testing"

	"example.com/adventOfCode/common"
	"example.com/adventOfCode/day9"
)

func TestTask1(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`
	res := day9.Task1(input)

	if res != 15 {
		t.Errorf("Expected 15 got %d", res)
	}
}

func TestTask1File(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day9.Task1(input)

	if res != 572 {
		t.Errorf("Expected 572 got %d", res)
	}
}

func TestTask2(t *testing.T) {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`
	res := day9.Task2(input)

	if res != 1134 {
		t.Errorf("Expected 1134 got %d", res)
	}
}

func TestTask2Test2(t *testing.T) {
	input := `9919199999
9911199999
1999999999
9999999999
1999999999
9999999999`
	res := day9.Task2(input)

	if res != 5 {
		t.Errorf("Expected 5 got %d", res)
	}
}

func TestTask2File(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day9.Task2(input)

	if res != 847044 {
		t.Errorf("Expected 847044 got %d", res)
	}
}
