package main

import (
	"fmt"
	"github.com/dusnm/AdventOfCode2023/pkg/day3"

	"github.com/dusnm/AdventOfCode2023/pkg/day1"
	"github.com/dusnm/AdventOfCode2023/pkg/day2"
)

func main() {
	solutions := map[string][]int{
		"day1": {
			day1.Puzzle1(day1.Input()),
			day1.Puzzle2(day1.Input()),
		},
		"day2": {
			day2.Puzzle1(day2.Input()),
			day2.Puzzle2(day2.Input()),
		},
		"day3": {
			day3.Puzzle1(day3.Input()),
			day3.Puzzle2(day3.Input()),
		},
	}

	fmt.Println(solutions)
}
