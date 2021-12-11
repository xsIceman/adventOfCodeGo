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

}
