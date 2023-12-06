package day6

import (
	"bufio"
	"strings"
	"unicode"
)

const (
	StateTime int = iota
	StateDistance
)

type (
	transition struct {
		state []rune
	}
)

func (t *transition) Append(c rune) {
	t.state = append(t.state, c)
}

func (t *transition) Next(c rune) bool {
	if len(t.state) == 0 {
		return false
	}

	lastState := t.state[len(t.state)-1]

	if unicode.IsDigit(c) && unicode.IsDigit(lastState) {
		return false
	}

	if c == ' ' && lastState == ' ' {
		return false
	}

	return true
}

func (t *transition) Reset() {
	t.state = make([]rune, 0)
}

func newScanner(input string) *bufio.Scanner {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

	return scanner
}

func NewMultipleRaces(input string) []Race {
	races := make([]Race, 4)
	scanner := newScanner(input)

	state := StateTime
	index := -1
	t := transition{}

	for scanner.Scan() {
		c := rune(scanner.Bytes()[0])

		switch state {
		case StateTime:
			if unicode.IsDigit(c) {
				races[index].Time = races[index].Time*10 + int(c-'0')
			}

			if c == ' ' && t.Next(c) {
				index++
			}

			t.Append(c)

			if c == '\n' {
				t.Reset()
				index = -1
				state = StateDistance
			}
		case StateDistance:
			if unicode.IsDigit(c) {
				races[index].Distance = races[index].Distance*10 + int(c-'0')
			}

			if c == ' ' && t.Next(c) {
				index++
			}

			t.Append(c)
		}
	}

	return races
}

func NewSingleRace(input string) Race {
	race := Race{}
	scanner := newScanner(input)

	state := StateTime
	for scanner.Scan() {
		c := rune(scanner.Bytes()[0])

		switch state {
		case StateTime:
			if unicode.IsDigit(c) {
				race.Time = race.Time*10 + int(c-'0')
			}

			if c == '\n' {
				state = StateDistance
			}
		case StateDistance:
			if unicode.IsDigit(c) {
				race.Distance = race.Distance*10 + int(c-'0')
			}
		}
	}

	return race
}
