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

func (h *Hand) DetermineKind(withJokers bool) int {
	counts := make(map[rune]int, len(h.Value))
	hasThreeOfAKind := false
	pairs, jokers := 0, 0

	for _, label := range h.Value {
		if withJokers && label == 'J' {
			jokers++
		} else {
			counts[label]++
		}
	}

	if withJokers {
		maximumKey := '0'
		maximumValue := 0
		for label, count := range counts {
			if count > maximumValue {
				maximumKey = label
				maximumValue = count
			}
		}

		counts[maximumKey] += jokers
	}

	for _, count := range counts {
		switch count {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			hasThreeOfAKind = true
		case 2:
			pairs++
		}
	}

	if hasThreeOfAKind && pairs == 1 {
		return FullHouse
	}

	if hasThreeOfAKind {
		return ThreeOfAKind
	}

	if pairs == 2 {
		return TwoPair
	}

	if pairs == 1 {
		return OnePair
	}

	return HighCard
}

func (g *Game) Sort(withJokers bool) {
	values := map[rune]int{
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

	if withJokers {
		values['J'] = 1
	}

	slices.SortFunc(g.Hands, func(a, b Hand) int {
		aKind := a.DetermineKind(withJokers)
		bKind := b.DetermineKind(withJokers)

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
