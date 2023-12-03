package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveB(input string) string {
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

	adjacent := make([][]Point, 0, len(symbols)*8)
	// Check 8 coordinates around the symbol.
	for _, symbol := range symbols {
		if symbol.Symbol == '*' {
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
			})
		}
	}

	total := 0
	for _, adj := range adjacent {
		uniqueValues := make(map[Value]struct{})
		for _, p := range adj {
			if value, ok := nums[p]; ok {
				uniqueValues[value] = struct{}{}
			}
		}
		if len(uniqueValues) == 2 {
			gearRatio := 1
			for k := range uniqueValues {
				gearRatio *= *k
			}
			total += gearRatio
		}
		uniqueValues = make(map[Value]struct{})
	}

	return fmt.Sprint(total)
}
