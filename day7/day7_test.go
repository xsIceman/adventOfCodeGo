package day7_test

import (
	"testing"

	"example.com/adventOfCode/common"
	"example.com/adventOfCode/day7"
)

func TestTask1(t *testing.T) {
	input := `16,1,2,0,4,2,7,1,2,14`
	res := day7.Task1(input)

	if res != 37 {
		t.Errorf("Expected to get 37 got %d", res)
	}
}

func TestTask1FileInput(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day7.Task1(input)

	if res != 323647 {
		t.Errorf("Expected to get 323647 got %d", res)
	}
}
