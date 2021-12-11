package day1

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func FindDept() int {
	reader, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	numbers, errParse := ReadInts(reader)
	if errParse != nil {
		log.Fatal(errParse)
	}

	lastnum := numbers[0]
	increase := 0

	for i := 1; i < len(numbers); i++ {
		if lastnum < numbers[i] {
			increase++
		}
		lastnum = numbers[i]
	}
	return increase
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}
