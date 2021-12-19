package day9

import (
	"sort"
	"strconv"
	"strings"
)

func Task1(input string) int {
	var heightMap [][]int

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	for _, l := range lines {
		var row []int
		for _, h := range strings.Split(l, "") {
			height, _ := strconv.Atoi(h)
			row = append(row, height)
		}
		heightMap = append(heightMap, row)
	}

	var lowpoints []int
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			current := heightMap[i][j]
			low := true
			if j+1 < len(heightMap[i]) {
				low = low && current < heightMap[i][j+1]
			}
			if j-1 >= 0 {
				low = low && current < heightMap[i][j-1]
			}
			if i-1 >= 0 {
				low = low && current < heightMap[i-1][j]
			}
			if i+1 < len(heightMap) {
				low = low && current < heightMap[i+1][j]
			}
			if low {
				lowpoints = append(lowpoints, current)
			}
		}
	}

	sum := 0
	for _, p := range lowpoints {
		sum = sum + 1 + p
	}
	return sum
}

func Task2(input string) int {
	var heightMap [][]int

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	for _, l := range lines {
		var row []int
		for _, h := range strings.Split(l, "") {
			height, _ := strconv.Atoi(h)
			row = append(row, height)
		}
		heightMap = append(heightMap, row)
	}

	var basinCount []int
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			current := heightMap[i][j]
			if current != 9 {
				basinCount = append(basinCount, countBasin(heightMap, i, j))
			}
		}
	}
	sort.Slice(basinCount, func(i, j int) bool {
		return basinCount[i] > basinCount[j]
	})

	return basinCount[0] * basinCount[1] * basinCount[2]
}

func countBasin(grid [][]int, i int, j int) int {
	if grid[i][j] != 9 {
		count := 1
		grid[i][j] = 9
		//down
		if i+1 < len(grid) {
			count = count + countBasin(grid, i+1, j)
		}
		//right
		if j+1 < len(grid[i]) {
			count = count + countBasin(grid, i, j+1)
		}
		//left
		if j-1 >= 0 {
			count = count + countBasin(grid, i, j-1)
		}
		//up
		if i-1 >= 0 {
			count = count + countBasin(grid, i-1, j)
		}
		return count
	}
	return 0
}
