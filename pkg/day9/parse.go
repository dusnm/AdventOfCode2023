package day9

import (
	"log"
	"strconv"
	"strings"
)

func New(input string) [][]int {
	lines := strings.Split(input, "\n")

	result := make([][]int, 0, len(lines))

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		values := strings.Split(line, " ")
		numbers := make([]int, 0, len(values))

		for _, number := range values {
			v, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			numbers = append(numbers, v)
		}

		result = append(result, numbers)
	}

	return result
}
