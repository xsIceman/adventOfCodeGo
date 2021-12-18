package day7

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func Task1(input string) int {

	var crabs []int
	inputNums := strings.Split(input, ",")
	for _, inputNum := range inputNums {
		age, _ := strconv.Atoi(inputNum)
		crabs = append(crabs, age)
	}

	sum := 0
	for _, v := range crabs {
		sum += v
	}

	sort.Ints(crabs)

	median := crabs[len(crabs)/2]

	minFuel := FindFuel(crabs, median)

	return minFuel

}

func FindFuel(crabs []int, pos int) int {
	fuel := 0
	for _, crab := range crabs {
		fuel += int(math.Abs(float64(crab - pos)))
	}

	return fuel
}
