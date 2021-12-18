package day8_test

import (
	"testing"

	"example.com/adventOfCode/common"
	"example.com/adventOfCode/day8"
)

func TestTask1(t *testing.T) {
	input := `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |	fdgacbe cefdb cefbgd gcbe
	edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |	fcgedb cgb dgebacf gc
	fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |	cg cg fdcagb cbg
	fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |	efabcd cedba gadfec cb
	aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |	gecf egdcabf bgf bfgea
	fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |	gebdcfa ecba ca fadegcb
	dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |	cefg dcbef fcge gbcadfe
	bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |	ed bcgafe cdgba cbgef
	egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |	gbdfcae bgc cg cgb
	gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |	fgae cfgab fg bagce`
	res := day8.Task1(input)

	if res != 26 {
		t.Errorf("Expected to get 26 got %d", res)
	}
}

func TestTask1FileInput(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day8.Task1(input)

	if res != 323647 {
		t.Errorf("Expected to get 323647 got %d", res)
	}
}
