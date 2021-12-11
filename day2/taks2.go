package day2

import (
	"strconv"
	"strings"
)

func Task2(inputs string) Position {
	pos := Position{Horizontal: 0,
		Depth: 0,
		Aim:   0}

	lines := strings.Split(strings.ReplaceAll(inputs, "\r\n", "\n"), "\n")

	for _, line := range lines {
		operation := strings.Split(line, " ")
		movement, _ := strconv.Atoi(operation[1])
		switch operation[0] {
		case "forward":
			pos.Horizontal += movement
			pos.Depth += (pos.Aim * movement)
		case "down":
			pos.Aim += movement
		case "up":
			pos.Aim -= movement
		}
	}
	return pos
}
