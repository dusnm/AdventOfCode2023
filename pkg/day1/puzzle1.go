package day1

import (
	"strings"
	"unicode"
)

func Puzzle1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		digits := []int{}
		for _, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, int(c-'0'))
			}
		}

		if size := len(digits); size != 0 {
			sum += digits[0]*10 + digits[size-1]
		}
	}

	return sum
}
