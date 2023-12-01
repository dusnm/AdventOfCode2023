package main

import (
	"fmt"

	"github.com/dusnm/AdventOfCode2023/pkg/day1"
)

func main() {
	solutions := map[string][]int{
		"day1": {
			day1.Puzzle1(day1.Input()),
			day1.Puzzle2(day1.Input()),
		},
	}

	fmt.Println(solutions)
}
