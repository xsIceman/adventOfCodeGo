package day5

import (
	"strconv"
	"strings"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func Task1(input string, toCheck int) int {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	max_col := 0
	max_row := 0
	var llLines []Line
	//73,572 -> 539,106
	for _, line := range lines {
		ll := Line{}
		points := strings.Split(strings.ReplaceAll(line, " -> ", ";"), ";")

		ll.x1, _ = strconv.Atoi(strings.Split(points[0], ",")[0])
		ll.y1, _ = strconv.Atoi(strings.Split(points[0], ",")[1])
		ll.x2, _ = strconv.Atoi(strings.Split(points[1], ",")[0])
		ll.y2, _ = strconv.Atoi(strings.Split(points[1], ",")[1])
		llLines = append(llLines, ll)
	}
	max_col = findMaxX(llLines)
	max_row = findMaxY(llLines)

	data := make([][]int, max_col+1)
	for i := range data {
		data[i] = make([]int, max_row+1)
	}

	for _, l := range llLines {
		if l.x1 == l.x2 {
			start := l.y1
			end := l.y2
			if l.y1 > l.y2 {
				start = l.y2
				end = l.y1
			}
			for i := start; i <= end; i++ {
				data[l.x1][i]++
			}
		}
		if l.y1 == l.y2 {
			start := l.x1
			end := l.x2
			if l.x1 > l.x2 {
				start = l.x2
				end = l.x1
			}
			for i := start; i <= end; i++ {
				data[i][l.y1]++
			}
		}
	}
	count := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] > toCheck {
				count++
			}
		}

	}

	return count
}

func Task2(input string, toCheck int) int {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	max_col := 0
	max_row := 0
	var llLines []Line
	//73,572 -> 539,106
	for _, line := range lines {
		ll := Line{}
		points := strings.Split(strings.ReplaceAll(line, " -> ", ";"), ";")

		ll.x1, _ = strconv.Atoi(strings.Split(points[0], ",")[0])
		ll.y1, _ = strconv.Atoi(strings.Split(points[0], ",")[1])
		ll.x2, _ = strconv.Atoi(strings.Split(points[1], ",")[0])
		ll.y2, _ = strconv.Atoi(strings.Split(points[1], ",")[1])
		llLines = append(llLines, ll)
	}
	max_col = findMaxX(llLines)
	max_row = findMaxY(llLines)

	data := make([][]int, max_col+1)
	for i := range data {
		data[i] = make([]int, max_row+1)
	}

	for _, l := range llLines {

		stepx := 1
		if l.x1 == l.x2 {
			stepx = 0
		}
		stepy := 1
		if l.y1 == l.y2 {
			stepy = 0
		}
		starty := l.y1
		endy := l.y2
		if l.y1 > l.y2 {
			stepy = -1
		}

		startx := l.x1
		endx := l.x2
		if l.x1 > l.x2 {
			stepx = -1
		}
		if stepx >= 0 && stepy >= 0 {
			for x, y := startx, starty; x <= endx && y <= endy; x, y = x+stepx, y+stepy {
				data[x][y]++
			}
		} else if stepx < 0 && stepy >= 0 {
			for x, y := startx, starty; x >= endx && y <= endy; x, y = x+stepx, y+stepy {
				data[x][y]++
			}
		} else if stepx >= 0 && stepy < 0 {
			for x, y := startx, starty; x <= endx && y >= endy; x, y = x+stepx, y+stepy {
				data[x][y]++
			}
		} else {
			for x, y := startx, starty; x >= endx && y >= endy; x, y = x+stepx, y+stepy {
				data[x][y]++
			}
		}

	}
	count := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] > toCheck {
				count++
			}
		}

	}

	return count
}

func findMaxX(lines []Line) int {
	max := 0
	for _, l := range lines {
		if l.x1 > max {
			max = l.x1
		}
		if l.x2 > max {
			max = l.x2
		}
	}
	return max
}

func findMaxY(lines []Line) int {
	max := 0
	for _, l := range lines {
		if l.y1 > max {
			max = l.y1
		}
		if l.y2 > max {
			max = l.y2
		}
	}
	return max
}
