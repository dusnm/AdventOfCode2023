package day4

import (
	"bufio"
	"strings"
	"unicode"
)

const (
	StateID int = iota
	StateWinningNumbers
	StateAvailableNumbers
	StateLineEnd
)

type (
	Game struct {
		ID        int
		Winning   []int
		Available []int
	}

	transition struct {
		state []rune
	}
)

func (t *transition) Append(c rune) {
	t.state = append(t.state, c)
}

func (t *transition) HasTransitioned(c rune) bool {
	if len(t.state) == 0 {
		return false
	}

	lastState := t.state[len(t.state)-1]

	// No transition
	if unicode.IsDigit(c) && unicode.IsDigit(lastState) {
		return false
	}

	// No transition
	if c == ' ' && lastState == ' ' {
		return false
	}

	return true
}

func (t *transition) Reset() {
	t.state = make([]rune, 0)
}

func (g *Game) SetID(c rune) {
	g.ID = g.ID*10 + int(c-'0')
}

func (g *Game) SetWinning(c rune, index int) {
	if len(g.Winning) <= index {
		g.Winning = append(g.Winning, int(c-'0'))
	} else {
		g.Winning[index] = g.Winning[index]*10 + int(c-'0')
	}
}

func (g *Game) SetAvailable(c rune, index int) {
	if len(g.Available) <= index {
		g.Available = append(g.Available, int(c-'0'))
	} else {
		g.Available[index] = g.Available[index]*10 + int(c-'0')
	}
}

func New(input string) []Game {
	// Parsing solution implemented using a Finite State Machine
	lines := strings.Count(input, "\n") + 1
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)

	// Initial state
	games := make([]Game, lines)
	lineCounter := 0
	winningTransition := transition{}
	availableTransition := transition{}
	winningCounter := 0
	availableCounter := 0
	state := StateID

	for scanner.Scan() {
		c := rune(scanner.Bytes()[0])
		switch state {
		case StateID:
			if unicode.IsDigit(c) {
				games[lineCounter].SetID(c)
			}

			if c == ':' {
				state = StateWinningNumbers
			}
		case StateWinningNumbers:
			if unicode.IsDigit(c) {
				games[lineCounter].SetWinning(c, winningCounter)
			}

			if c == ' ' && winningTransition.HasTransitioned(c) {
				winningCounter++
			}

			winningTransition.Append(c)

			if c == '|' {
				state = StateAvailableNumbers
			}
		case StateAvailableNumbers:
			if unicode.IsDigit(c) {
				games[lineCounter].SetAvailable(c, availableCounter)
			}

			if c == ' ' && availableTransition.HasTransitioned(c) {
				availableCounter++
			}

			availableTransition.Append(c)

			if c == '\n' {
				state = StateLineEnd
			}
		case StateLineEnd:
			lineCounter++
			winningCounter = 0
			availableCounter = 0
			winningTransition.Reset()
			availableTransition.Reset()
			state = StateID
		}
	}

	return games
}
