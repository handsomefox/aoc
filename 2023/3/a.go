package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveA(input string) string {
	var (
		sc      = bufio.NewScanner(strings.NewReader(input))
		symbols = make([]Symbol, 0)
		nums    = make(map[Point]Value)
	)
	for y := 0; sc.Scan(); y++ {
		txt := sc.Text()
		// Get symbols
		symbols = append(symbols, parseSymbols(txt, y)...)
		// Get numbers
		for _, num := range parseNums(txt, y) {
			for _, p := range num.Points {
				nums[p] = num.Value
			}
		}
	}
	// Only store unique values (since the symbol might encounter the same number multiple times when checking around itself)
	uniqueValues := make(map[Value]struct{})

	adjacent := make([]Point, 0, len(symbols)*8)
	// Check 8 coordinates around the symbol.
	for _, symbol := range symbols {
		adjacent = append(adjacent, []Point{
			// Left side of the symbol
			{X: symbol.X - 1, Y: symbol.Y - 1},
			{X: symbol.X - 1, Y: symbol.Y},
			{X: symbol.X - 1, Y: symbol.Y + 1},

			// Above and below the symbol
			{X: symbol.X, Y: symbol.Y - 1},
			{X: symbol.X, Y: symbol.Y + 1},

			// Right side of the symbol
			{X: symbol.X + 1, Y: symbol.Y - 1},
			{X: symbol.X + 1, Y: symbol.Y},
			{X: symbol.X + 1, Y: symbol.Y + 1},
		}...)
	}

	for _, p := range adjacent {
		if value, ok := nums[p]; ok {
			uniqueValues[value] = struct{}{}
		}
	}

	sum := 0
	for k := range uniqueValues {
		sum += *k
	}

	return fmt.Sprint(sum)
}
