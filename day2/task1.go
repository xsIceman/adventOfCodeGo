package day2

import (
	"strconv"
	"strings"
)

type Position struct {
	Horizontal int
	Depth      int
	Aim        int
}

func Task1(inputs string) Position {
	pos := Position{Horizontal: 0,
		Depth: 0}

	lines := strings.Split(strings.ReplaceAll(inputs, "\r\n", "\n"), "\n")

	for _, line := range lines {
		operation := strings.Split(line, " ")
		movement, _ := strconv.Atoi(operation[1])
		switch operation[0] {
		case "forward":
			pos.Horizontal += movement
		case "down":
			pos.Depth += movement
		case "up":
			pos.Depth -= movement
		}
	}
	return pos
}
