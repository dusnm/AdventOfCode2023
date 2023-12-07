package day7

import (
	"slices"
)

const (
	HighCard int = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type (
	Hand struct {
		Bid   int
		Value string
	}

	Game struct {
		Hands []Hand
	}
)

var (
	values = map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
)

func (h *Hand) DetermineKind() int {
	hand := []rune(h.Value)
	clusters := make([][]rune, len(hand))

	slices.SortFunc(hand, func(a, b rune) int {
		v1, v2 := values[a], values[b]

		if v1 == v2 {
			return 0
		}

		if v1 < v2 {
			return -1
		}

		return 0
	})

	index := 0
	clusters[index] = []rune{hand[0]}
	for i := 1; i < len(hand); i++ {
		current := hand[i]
		previous := hand[i-1]

		if current == previous {
			clusters[index] = append(clusters[index], current)
		} else {
			index++
			clusters[index] = append(clusters[index], current)
		}
	}

	c1 := clusters[0]
	c2 := clusters[1]
	c3 := clusters[2]
	c4 := clusters[3]
	c5 := clusters[4]

	if len(c1) == 5 {
		return FiveOfAKind
	}

	if len(c1) == 4 || len(c2) == 4 {
		return FourOfAKind
	}

	if len(c1) == 3 && len(c2) == 2 || len(c2) == 3 && len(c1) == 2 {
		return FullHouse
	}

	if len(c1) == 3 || len(c2) == 3 || len(c3) == 3 {
		return ThreeOfAKind
	}

	if len(c1) == 2 && len(c2) == 2 || len(c1) == 2 && len(c3) == 2 || len(c2) == 2 && len(c3) == 2 {
		return TwoPair
	}

	if len(c1) == 2 || len(c2) == 2 || len(c3) == 2 || len(c4) == 2 || len(c5) == 2 {
		return OnePair
	}

	return HighCard
}

func (g *Game) Sort() {
	slices.SortFunc(g.Hands, func(a, b Hand) int {
		aKind := a.DetermineKind()
		bKind := b.DetermineKind()

		if aKind < bKind {
			return -1
		}

		if bKind < aKind {
			return 1
		}

		for k := range a.Value {
			l1 := rune(a.Value[k])
			l2 := rune(b.Value[k])

			if v1, v2 := values[l1], values[l2]; v1 != v2 {
				if v1 < v2 {
					return -1
				}

				return 1
			}
		}

		return 0
	})
}
