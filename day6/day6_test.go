package day6_test

import (
	"testing"

	"example.com/adventOfCode/common"
	"example.com/adventOfCode/day6"
)

func TestTask1(t *testing.T) {
	input := `3,4,3,1,2`
	res := day6.Task1(input, 18)
	if res != 26 {
		t.Errorf("Expected 26 points got %d", res)
	}
	res = day6.Task1(input, 80)
	if res != 5934 {
		t.Errorf("Expected 5934 points got %d", res)
	}
}

func TestTask1File(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day6.Task1(input, 18)
	if res != 1635 {
		t.Errorf("Expected 1635 points got %d", res)
	}
	res = day6.Task1(input, 80)
	if res != 360761 {
		t.Errorf("Expected 360761 points got %d", res)
	}
}
