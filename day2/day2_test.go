package day2_test

import (
	"testing"

	"example.com/adventOfCode/day2"
)

func TestDay2Task1(t *testing.T) {
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	pos := day2.Task1(input)

	if pos.Depth != 10 {
		t.Errorf("depth expected to be 10 got %d", pos.Depth)
	}
	if pos.Horizontal != 15 {
		t.Errorf("Horizontal expected to be 15 got %d", pos.Horizontal)
	}

}

func TestDay2Task1File(t *testing.T) {
	input := day2.ReadFile("input.txt")
	pos := day2.Task1(input)

	res := pos.Depth * pos.Horizontal
	if res != 1694130 {
		t.Errorf("res expected to be 10 got %d", res)
	}
	if pos.Depth != 894 {
		t.Errorf("depth expected to be 10 got %d", pos.Depth)
	}
	if pos.Horizontal != 1895 {
		t.Errorf("Horizontal expected to be 15 got %d", pos.Horizontal)
	}
}

func TestDay2Task2(t *testing.T) {
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	pos := day2.Task2(input)

	if pos.Depth != 60 {
		t.Errorf("depth expected to be 10 got %d", pos.Depth)
	}
	if pos.Horizontal != 15 {
		t.Errorf("Horizontal expected to be 15 got %d", pos.Horizontal)
	}

}

func TestDay2Task2File(t *testing.T) {
	input := day2.ReadFile("input.txt")
	pos := day2.Task2(input)

	res := pos.Depth * pos.Horizontal
	if res != 1698850445 {
		t.Errorf("res expected to be 1698850445 got %d", res)
	}
	if pos.Depth != 896491 {
		t.Errorf("depth expected to be 10 got %d", pos.Depth)
	}
	if pos.Horizontal != 1895 {
		t.Errorf("Horizontal expected to be 15 got %d", pos.Horizontal)
	}
}
