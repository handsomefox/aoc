package main

import (
	"fmt"

	"aoc/shared"
)

const (
	year = 2025
	day  = 3
)

func main() {
	input := shared.MustReadInput(year, day)

	fmt.Println("Part A:", PartA(input))
	fmt.Println("Part B:", PartB(input))
}

func PartA(input string) string {
	return solve(input, 2)
}

func PartB(input string) string {
	return solve(input, 12)
}

func solve(input string, charsNeeded int) string {
	total := 0

	for line := range shared.Lines(input) {
		totalStr := ""
		offset := 0

		for c := range charsNeeded {
			max := 0
			for i := offset; i < len(line); i++ {
				value := shared.Btoi(line[i])
				if value > max {
					max = value
					offset = i + 1
				}
				remaining := len(line) - i
				if remaining <= charsNeeded-c {
					break
				}
			}
			totalStr += shared.String(max)
		}

		total += shared.MustParseInt(totalStr)
	}

	return shared.String(total)
}
