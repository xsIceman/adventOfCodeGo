package day5_test

import (
	"testing"

	"example.com/adventOfCode/common"
	"example.com/adventOfCode/day5"
)

func TestTask1(t *testing.T) {
	input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	res := day5.Task1(input, 1)
	if res != 5 {
		t.Errorf("Expected 5 points got %d", res)
	}
}

func TestTask1MoreThan1(t *testing.T) {
	input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	res := day5.Task1(input, 0)
	if res != 21 {
		t.Errorf("Expected 21 points got %d", res)
	}
}

func TestTask1File(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day5.Task1(input, 1)
	if res != 6461 {
		t.Errorf("Expected 6461 points got %d", res)
	}
}

func TestTask2(t *testing.T) {
	input := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
	res := day5.Task2(input, 1)
	if res != 12 {
		t.Errorf("Expected 12 points got %d", res)
	}
}

func TestTask2File(t *testing.T) {
	input := common.ReadFile("input.txt")

	res := day5.Task2(input, 1)
	if res != 18065 {
		t.Errorf("Expected 18065 points got %d", res)
	}
}
