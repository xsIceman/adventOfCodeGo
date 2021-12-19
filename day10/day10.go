package day10

import (
	"fmt"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func Task1(input string) int {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	errorSum := 0
	for _, l := range lines {
		lineStack := stack.New()
		for ix, s := range strings.Split(l, "") {
			c := rune(s[0])

			if c == '(' || c == '{' || c == '[' || c == '<' {
				lineStack.Push(c)
			} else {
				peek := int(lineStack.Peek().(rune))
				if peek+1 == int(c) || peek+2 == int(c) {
					lineStack.Pop()
				} else {
					fmt.Print(lineStack.Peek())
					fmt.Print(ix)
					if c == ')' {
						errorSum = errorSum + 3
					}
					if c == ']' {
						errorSum = errorSum + 57
					}
					if c == '}' {

						errorSum = errorSum + 1197
					}
					if c == '>' {
						errorSum = errorSum + 25137
					}
					break
				}
			}
		}
	}
	return errorSum
}
