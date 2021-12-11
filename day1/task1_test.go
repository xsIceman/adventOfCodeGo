package day1_test

import (
	"testing"

	"example.com/adventOfCode/day1"
)

func TestTask1(t *testing.T) {
	increases := day1.FindDept()

	if increases != 1184 {
		t.Errorf("expected to get 1884 got %d", increases)
	}
}

func TestTask2(t *testing.T) {
	numbers := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	increases := day1.Task2(numbers)

	if increases != 5 {
		t.Errorf("expected to get 5 got %d", increases)
	}
}

func TestTask2WithFile(t *testing.T) {
	increases := day1.Task2(nil)

	if increases != 1158 {
		t.Errorf("expected to get 5 got %d", increases)
	}
}
