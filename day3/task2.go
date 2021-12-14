package day3

import (
	"strconv"
	"strings"
)

func Task2(input string) Diagnostic {
	diag := Diagnostic{}

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	chOxy := make(chan int)
	chCo2 := make(chan int)
	go findOxy(lines, chOxy)
	go findCo2(lines, chCo2)

	diag.CO2Scrubber = <-chCo2
	diag.OxygenGenerator = <-chOxy

	diag.LifeSupport = diag.CO2Scrubber * diag.OxygenGenerator

	return diag
}

func findCo2(lines []string, chco2 chan int) {
	linesPrevious := lines

	for i := 0; i < len(lines[0]); i++ {
		linesThisRount := linesPrevious
		bitCount := getBits(linesThisRount)
		charToCheck := byte('1')

		if float64(bitCount[i]) >= (float64(len(linesThisRount)) / 2.0) {
			charToCheck = '0'
		} else {
			charToCheck = '1'
		}

		linesPrevious = []string{}
		for _, line := range linesThisRount {
			if line[i] == charToCheck {
				linesPrevious = append(linesPrevious, line)
			}
		}
		if len(linesPrevious) == 1 {
			res, _ := strconv.ParseInt(linesPrevious[0], 2, 32)
			chco2 <- int(res)
			break
		}
	}
	chco2 <- -1
}

func findOxy(lines []string, chOxy chan int) {
	linesPrevious := lines
	for i := 0; i < len(lines[0]); i++ {
		linesThisRount := linesPrevious
		bitCount := getBits(linesThisRount)
		charToCheck := byte('1')

		if float64(bitCount[i]) >= (float64(len(linesThisRount)) / 2.0) {
			charToCheck = '1'
		} else {
			charToCheck = '0'
		}

		linesPrevious = []string{}
		for _, line := range linesThisRount {
			if line[i] == charToCheck {
				linesPrevious = append(linesPrevious, line)
			}
		}
		if len(linesPrevious) == 1 {
			res, _ := strconv.ParseInt(linesPrevious[0], 2, 32)
			chOxy <- int(res)
		}
	}
	chOxy <- -1
}

func getBits(lines []string) []int {
	bitCount := make([]int, len(lines[0]))
	for _, line := range lines {
		bits := []rune(line)
		for i := 0; i < len(bits); i++ {
			if bits[i] == '1' {
				bitCount[i]++
			}
		}
	}
	return bitCount
}
