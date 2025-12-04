package main

import (
	"fmt"
	"strings"

	"aoc/shared"
)

const (
	year = 2025
	day  = 4
)

func main() {
	input := shared.MustReadInput(year, day)

	fmt.Println("Part A:", PartA(input))
	fmt.Println("Part B:", PartB(input))
}

func PartA(input string) string {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	total := 0
	for x := range len(lines[0]) {
		for y := range lines {
			if lines[y][x] == '.' {
				continue
			}
			position := position{x, y}
			xBound, yBound := len(lines[0]), len(lines)
			adjacent := position.adjacentPositions(xBound, yBound)
			rolls := countRolls(lines, adjacent)
			if rolls < 4 {
				total++
			}
		}
	}
	return shared.String(total)
}

func PartB(input string) string {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	totalRemoved := 0
	for {
		var remove []position
		for x := range len(lines[0]) {
			for y := range lines {
				if lines[y][x] == '.' {
					continue
				}
				position := position{x, y}
				xBound, yBound := len(lines[0]), len(lines)
				adjacent := position.adjacentPositions(xBound, yBound)
				rolls := countRolls(lines, adjacent)
				if rolls < 4 {
					remove = append(remove, position)
				}
			}
		}
		if len(remove) == 0 {
			break
		}
		for _, rm := range remove {
			lines[rm.y] = shared.ReplaceAtStringIndex(lines[rm.y], rm.x, '.')
		}
		totalRemoved += len(remove)
	}
	return shared.String(totalRemoved)
}

type position struct {
	x, y int
}

func (p position) adjacentPositions(xBound, yBound int) []position {
	// positions := []pos{
	// 	{x-1,y-1}, {x, y-1}, {x+1, y-1},
	// 	{x-1, y}, /* curr, */ {x+1, y},
	// 	{x-1, y+1}, {x, y+1}, {x+1, y+1},
	// }
	x, y := p.x, p.y
	adjacent := [8]position{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
		{x - 1, y},
		{x + 1, y},
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}

	return shared.Filter(adjacent[:], func(value position) bool {
		return value.inbounds(xBound, yBound)
	})
}

func (p position) inbounds(xBound, yBound int) bool {
	return p.x >= 0 && p.x < xBound && p.y >= 0 && p.y < yBound
}

func countRolls(lines []string, position []position) int {
	total := 0
	for _, pos := range position {
		if lines[pos.y][pos.x] == '@' {
			total++
		}
	}
	return total
}
