package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dusnm/AdventOfCode2023/pkg/day1"
	"github.com/dusnm/AdventOfCode2023/pkg/day2"
	"github.com/dusnm/AdventOfCode2023/pkg/day3"
	"github.com/dusnm/AdventOfCode2023/pkg/day4"
	"github.com/dusnm/AdventOfCode2023/pkg/day5"
	"github.com/dusnm/AdventOfCode2023/pkg/day6"
	"github.com/dusnm/AdventOfCode2023/pkg/day7"
)

var dayFlag int

func init() {
	flag.IntVar(
		&dayFlag,
		"day",
		0,
		"Day number. Runs the corresponding solution for the day.",
	)

	flag.Parse()
}

type (
	SolutionFunc func() []int
)

func main() {
	solutions := []SolutionFunc{
		func() []int {
			return []int{
				day1.Puzzle1(day1.Input()),
				day1.Puzzle2(day1.Input()),
			}
		},
		func() []int {
			return []int{
				day2.Puzzle1(day2.Input()),
				day2.Puzzle2(day2.Input()),
			}
		},
		func() []int {
			return []int{
				day3.Puzzle1(day3.Input()),
				day3.Puzzle2(day3.Input()),
			}
		},
		func() []int {
			return []int{
				day4.Puzzle1(day4.Input()),
				day4.Puzzle2(day4.Input()),
			}
		},
		func() []int {
			return []int{
				day5.Puzzle1(day5.Input()),
				day5.Puzzle2(day5.Input()),
			}
		},
		func() []int {
			return []int{
				day6.Puzzle1(day6.Input()),
				day6.Puzzle2(day6.Input()),
			}
		},
		func() []int {
			return []int{
				day7.Puzzle1(day7.Input()),
				day7.Puzzle2(day7.Input()),
			}
		},
	}

	if len(solutions) < dayFlag {
		log.Fatalf("No solution for day %d\n", dayFlag)
	}

	fmt.Println(solutions[dayFlag-1]())
}
