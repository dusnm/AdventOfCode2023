package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type (
	Matrix struct {
		Data   [][]rune
		Height int
		Width  int
	}

	Adjacent struct {
		Value rune
		X     int
		Y     int
	}
)

func (a Adjacent) String() string {
	return fmt.Sprintf("%dx%d", a.X, a.Y)
}

func load(input string) *Matrix {
	lines := strings.Split(input, "\n")
	data := make([][]rune, 0, len(lines))

	for _, line := range lines {
		inner := make([]rune, 0, len(line))
		for _, c := range line {
			inner = append(inner, c)
		}

		data = append(data, inner)
	}

	return &Matrix{
		Data:   data,
		Height: len(lines),
		Width:  len(data[0]),
	}
}

func (m *Matrix) IsBounded(row, column int) bool {
	if row >= 0 && row < m.Height && column >= 0 && column < m.Width {
		return true
	}

	return false
}

func (m *Matrix) GetAdjacencyList(row, column int) []Adjacent {
	adjacencyList := make([]Adjacent, 0, 8)
	if m.IsBounded(row+1, column+1) {
		adjacencyList = append(adjacencyList, Adjacent{
			Value: m.Data[row+1][column+1],
			X:     row + 1,
			Y:     column + 1,
		})
	}

	if m.IsBounded(row+1, column) {
		adjacencyList = append(adjacencyList, Adjacent{
			Value: m.Data[row+1][column],
			X:     row + 1,
			Y:     column,
		})
	}

	if m.IsBounded(row+1, column-1) {
		adjacencyList = append(adjacencyList, Adjacent{
			Value: m.Data[row+1][column-1],
			X:     row + 1,
			Y:     column - 1,
		})
	}

	if m.IsBounded(row, column+1) {
		adjacencyList = append(adjacencyList, Adjacent{
			Value: m.Data[row][column+1],
			X:     row,
			Y:     column + 1,
		})
	}

	if m.IsBounded(row, column-1) {
		adjacencyList = append(adjacencyList, Adjacent{
			Value: m.Data[row][column-1],
			X:     row,
			Y:     column - 1,
		})
	}

	if m.IsBounded(row-1, column+1) {
		adjacencyList = append(adjacencyList, Adjacent{
			Value: m.Data[row-1][column+1],
			X:     row - 1,
			Y:     column + 1,
		})
	}

	if m.IsBounded(row-1, column) {
		adjacencyList = append(adjacencyList, Adjacent{
			Value: m.Data[row-1][column],
			X:     row - 1,
			Y:     column,
		})
	}

	if m.IsBounded(row-1, column-1) {
		adjacencyList = append(adjacencyList, Adjacent{
			Value: m.Data[row-1][column-1],
			X:     row - 1,
			Y:     column - 1,
		})
	}

	return adjacencyList
}

func (m *Matrix) FindPartNumbers() []PartNumber {
	parts := make([]PartNumber, 0)

	for row, line := range m.Data {
		adjacency := make([]Adjacent, 0)
		builder := strings.Builder{}
		for column, c := range line {
			if unicode.IsDigit(c) {
				builder.WriteRune(c)
				adjacency = append(adjacency, m.GetAdjacencyList(row, column)...)

				if column == len(line)-1 {
					number := builder.String()
					value, _ := strconv.Atoi(number)
					parts = append(parts, PartNumber{
						Value:     value,
						Adjacency: adjacency,
						Length:    len(number),
					})

					builder.Reset()
					adjacency = make([]Adjacent, 0)
				}
			} else {
				if builder.Len() != 0 {
					number := builder.String()
					value, _ := strconv.Atoi(number)
					parts = append(parts, PartNumber{
						Value:     value,
						Adjacency: adjacency,
						Length:    len(number),
					})

					builder.Reset()
					adjacency = make([]Adjacent, 0)
				}
			}
		}
	}

	filteredParts := make([]PartNumber, 0, len(parts))
	for _, part := range parts {
		adjacency := part.Adjacency

		for _, c := range adjacency {
			if c.Value == '.' {
				continue
			}

			if !unicode.IsDigit(c.Value) {
				filteredParts = append(filteredParts, part)
				break
			}
		}
	}

	return filteredParts
}
