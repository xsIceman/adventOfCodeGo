package day3

import (
	"strconv"
	"strings"
)

type Diagnostic struct {
	Gamma            int
	Epsilon          int
	PowerConsumption int
	OxygenGenerator  int
	CO2Scrubber      int
	LifeSupport      int
}

func Task1(input string) Diagnostic {
	diag := Diagnostic{Gamma: 0}

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	bitCount := make([]int, len(lines[0]))

	for _, line := range lines {
		bits := []rune(line)
		for i := 0; i < len(bits); i++ {
			if bits[i] == '1' {
				bitCount[i]++
			}
		}
	}

	gamma := ""
	epsilon := ""

	for _, bit := range bitCount {
		if bit > (len(lines) / 2) {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}

	}

	res, _ := strconv.ParseInt(gamma, 2, 32)
	diag.Gamma = int(res)
	res, _ = strconv.ParseInt(epsilon, 2, 32)
	diag.Epsilon = int(res)

	diag.PowerConsumption = diag.Gamma * diag.Epsilon

	return diag
}
