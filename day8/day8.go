package day8

import (
	"strings"
)

func Task1(input string) int {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	var outputLines []string
	count := 0
	for _, line := range lines {
		output := strings.Split(line, "|")[1]
		outputLines = append(outputLines, output)
		output = strings.TrimSpace(output)
		for _, num := range strings.Split(output, " ") {
			if len(num) == 2 || len(num) == 4 || len(num) == 3 || len(num) == 7 {
				count++
			}
		}

	}
	return count

}
