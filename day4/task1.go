package day4

import (
	"strconv"
	"strings"
)

func Task1(input string) int {
	var boards [][5][5]int
	var draw_numbers []int

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	for _, num := range strings.Split(lines[0], ",") {
		intNum, _ := strconv.Atoi(num)
		draw_numbers = append(draw_numbers, intNum)
	}

	for i := 2; i < len(lines); i += 6 {
		var board [5][5]int
		for j := i; j < i+5; j++ {
			line := lines[j]
			numbers := strings.Split(strings.ReplaceAll(strings.TrimSpace(line), "  ", " "), " ")
			for ix, num := range numbers {
				board[j-i][ix], _ = strconv.Atoi(num)
			}
		}
		boards = append(boards, board)
	}

	for count, draw := range draw_numbers {

		//update boards
		for b := 0; b < len(boards); b++ {
			for x := 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					if boards[b][x][y] == draw {
						boards[b][x][y] = -1
					}
				}
			}
		}
		if count > 3 {
			for i := 0; i < len(boards); i++ {
				b := boards[i]
				if CheckBoardForBingo(b) {
					boardSum := SumBoard(b)
					return boardSum * draw
				}
			}
			//check boards
		}
	}

	return -1
}

func Task2(input string) int {
	var boards [][5][5]int
	var draw_numbers []int

	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	for _, num := range strings.Split(lines[0], ",") {
		intNum, _ := strconv.Atoi(num)
		draw_numbers = append(draw_numbers, intNum)
	}

	for i := 2; i < len(lines); i += 6 {
		var board [5][5]int
		for j := i; j < i+5; j++ {
			line := lines[j]
			numbers := strings.Split(strings.ReplaceAll(strings.TrimSpace(line), "  ", " "), " ")
			for ix, num := range numbers {
				board[j-i][ix], _ = strconv.Atoi(num)
			}
		}
		boards = append(boards, board)
	}

	for count, draw := range draw_numbers {

		//update boards
		for b := 0; b < len(boards); b++ {
			for x := 0; x < 5; x++ {
				for y := 0; y < 5; y++ {
					if boards[b][x][y] == draw {
						boards[b][x][y] = -1
					}
				}
			}
		}
		if count > 3 {
			for i := 0; i < len(boards); i++ {
				b := boards[i]
				if CheckBoardForBingo(b) {

					if len(boards) == 1 {
						boardSum := SumBoard(b)
						return boardSum * draw
					}
					boards = remove(boards, i)
				}
			}
			//check boards
		}
	}

	return -1
}

func remove(s [][5][5]int, i int) [][5][5]int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func SumBoard(board [5][5]int) int {
	sum := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if board[y][x] != -1 {
				sum += board[y][x]
			}
		}
	}
	return sum
}

func CheckBoardForBingo(board [5][5]int) bool {
	for x := 0; x < 5; x++ {
		if CheckRow(board, x) {
			return true
		}
		if CheckColumn(board, x) {
			return true
		}
	}
	return false
}

func CheckColumn(board [5][5]int, x int) bool {
	for y := 0; y < 5; y++ {
		if board[y][x] != -1 {
			return false
		}
	}
	return true
}

func CheckRow(board [5][5]int, x int) bool {
	for y := 0; y < 5; y++ {
		if board[x][y] != -1 {
			return false
		}
	}
	return true
}
