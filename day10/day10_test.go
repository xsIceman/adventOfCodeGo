package day10_test

import (
	"testing"

	"example.com/adventOfCode/common"
	"example.com/adventOfCode/day10"
)

func TestTask1(t *testing.T) {
	input := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`
	res := day10.Task1(input)

	if res != 26397 {
		t.Errorf("Expected 26397 got %d", res)
	}
}

func TestTask1File(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day10.Task1(input)

	if res != 193275 {
		t.Errorf("Expected 193275 got %d", res)
	}
}
