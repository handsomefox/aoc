package main

import (
	"fmt"
	"strings"

	"aoc/shared"
)

const (
	year = 2025
	day  = 7
)

func main() {
	input := shared.MustReadInput(year, day)

	fmt.Println("Part A:", PartA(input))
	fmt.Println("Part B:", PartB(input))
}

func PartA(input string) string {
	grid := shared.LinesSlice(input)

	var startX, startY int
	for y, row := range grid {
		if x := strings.Index(row, "S"); x != -1 {
			startX, startY = x, y
			break
		}
	}

	beams := map[int]struct{}{startX: {}}
	splits := 0

	for y := startY + 1; y < len(grid); y++ {
		next := map[int]struct{}{}

		for x := range beams {
			if grid[y][x] == '^' {
				splits++

				if x-1 >= 0 {
					next[x-1] = struct{}{}
				}
				if x+1 < len(grid[y]) {
					next[x+1] = struct{}{}
				}
				continue
			}

			next[x] = struct{}{}
		}

		beams = next
	}

	return shared.String(splits)
}

func PartB(input string) string {
	grid := shared.LinesSlice(input)

	var startX, startY int
	for y, row := range grid {
		if x := strings.Index(row, "S"); x != -1 {
			startX, startY = x, y
			break
		}
	}

	beams := map[int]int{startX: 1}
	exited := 0

	for y := startY + 1; y < len(grid); y++ {
		next := map[int]int{}

		for x, count := range beams {
			if grid[y][x] == '^' {
				if x-1 >= 0 {
					next[x-1] += count
				} else {
					exited += count
				}

				if x+1 < len(grid[y]) {
					next[x+1] += count
				} else {
					exited += count
				}
				continue
			}

			if y+1 < len(grid) {
				next[x] += count
			} else {
				exited += count
			}
		}

		beams = next
	}

	for _, c := range beams {
		exited += c
	}

	return shared.String(exited)
}
