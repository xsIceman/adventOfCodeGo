package day6

import (
	"strconv"
	"strings"
)

func Task1(input string, rounds int) int {
	var fish []int
	inputNums := strings.Split(input, ",")
	for _, inputNum := range inputNums {
		age, _ := strconv.Atoi(inputNum)
		fish = append(fish, age)
	}

	for r := 0; r < rounds; r++ {

		for i := len(fish) - 1; i >= 0; i-- {
			if fish[i] == 0 {
				fish = append(fish, 8)
				fish[i] = 6
			} else {
				fish[i]--
			}
		}
	}
	return len(fish)
}
