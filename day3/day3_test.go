package day3_test

import (
	"testing"

	"example.com/adventOfCode/common"
	"example.com/adventOfCode/day3"
)

func TestDay3Task1(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	res := day3.Task1(input)

	if res.Epsilon != 9 {
		t.Errorf("Expected Epsilon to be 9 go %d", res.Epsilon)
	}

	if res.Gamma != 22 {
		t.Errorf("Expected Gamma to be 22 go %d", res.Gamma)
	}

	if res.PowerConsumption != 198 {
		t.Errorf("Expected PowerConsumption to be 198 go %d", res.PowerConsumption)
	}
}

func TestDay3Task1File(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day3.Task1(input)

	if res.Epsilon != 2758 {
		t.Errorf("Expected Epsilon to be 2758 go %d", res.Epsilon)
	}

	if res.Gamma != 1337 {
		t.Errorf("Expected Gamma to be 1337 go %d", res.Gamma)
	}

	if res.PowerConsumption != 3687446 {
		t.Errorf("Expected PowerConsumption to be 3687446 go %d", res.PowerConsumption)
	}
}

func TestDay3Task2(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	res := day3.Task2(input)

	if res.OxygenGenerator != 23 {
		t.Errorf("Expected OxygenGenerator to be 23 %d", res.OxygenGenerator)
	}

	if res.CO2Scrubber != 10 {
		t.Errorf("Expected CO2Scrubber to be 10 go %d", res.CO2Scrubber)
	}

	if res.LifeSupport != 230 {
		t.Errorf("Expected LifeSupport to be 230 go %d", res.PowerConsumption)
	}
}

func TestDay3Task2File(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day3.Task2(input)

	if res.OxygenGenerator != 1599 {
		t.Errorf("Expected OxygenGenerator to be 1599 %d", res.OxygenGenerator)
	}

	if res.CO2Scrubber != 2756 {
		t.Errorf("Expected CO2Scrubber to be 2756 go %d", res.CO2Scrubber)
	}

	if res.LifeSupport != 4406844 {
		t.Errorf("Expected LifeSupport to be 4406844 go %d", res.LifeSupport)
	}
}
