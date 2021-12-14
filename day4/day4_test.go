package day4_test

import (
	"testing"

	"example.com/adventOfCode/common"
	"example.com/adventOfCode/day4"
)

func TestTask1(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`
	res := day4.Task1(input)
	if res != 4512 {
		t.Errorf("bingo result should be 4512 was  %d", res)
	}
}

func TestTask2(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`
	res := day4.Task2(input)
	if res != 1924 {
		t.Errorf("bingo result should be 1924 was  %d", res)
	}
}

func TestTask1Column(t *testing.T) {
	input := `22,8,21,6,1,2

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`
	res := day4.Task1(input)
	if res != 242 {
		t.Errorf("bingo result should be 242 was  %d", res)
	}
}

func TestTask1Row2(t *testing.T) {
	input := `9,18,13,17,5

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`
	res := day4.Task1(input)
	if res != 1310 {
		t.Errorf("bingo result should be 1310 was  %d", res)
	}
}

func TestTask1Column2(t *testing.T) {
	input := `15,18,8,11,21

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`
	res := day4.Task1(input)
	if res != 5271 {
		t.Errorf("bingo result should be 5271 was  %d", res)
	}
}

func TestTask1File(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day4.Task1(input)
	if res != 22680 {
		t.Errorf("bingo result should be 22680 was  %d", res)
	}
}

func TestTask2File(t *testing.T) {
	input := common.ReadFile("input.txt")
	res := day4.Task2(input)
	if res != 16168 {
		t.Errorf("bingo result should be 16168 was  %d", res)
	}
}
