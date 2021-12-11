package day1

import (
	"log"
	"os"
)

func SumArray(array []int) int {
	res := 0

	for i := 0; i < len(array); i++ {
		res += array[i]
	}
	return res
}

func Task2(inputnumbers []int) int {
	var numbers []int
	if inputnumbers == nil {
		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatal(err)
		}
		numbers, err = ReadInts(file)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		numbers = inputnumbers
	}

	increases := 0

	part1 := []int{numbers[0], numbers[1], numbers[2]}
	part2 := []int{numbers[1], numbers[2], numbers[3]}

	for i, j := 2, 3; j < len(numbers); i, j = i+1, j+1 {
		part1[i%3] = numbers[i]
		part2[i%3] = numbers[j]
		if SumArray(part2) > SumArray(part1) {
			increases++
		}
	}
	return increases
}
