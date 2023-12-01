package day1

import (
	"strings"
)

func Puzzle2(input string) int {
	// Sometimes you can get lines containing strings that
	// contain two adjacent string digits like twone. To
	// deal with this, we put our numeric digits between the first and last letter
	// to make sure we count both 1 and 2
	replacements := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	// This is also the reason why we can't replace them
	// all at once with something like strigs.Replacer
	for k, v := range replacements {
		input = strings.ReplaceAll(input, k, v)
	}

	return Puzzle1(input)
}
