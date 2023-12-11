package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

func Solve(input string, scale int) string {
	grid := Lines(input)

	emptyRows := []int{}
	for r, row := range grid {
		if All([]byte(row), func(b byte) bool { return b == '.' }) {
			emptyRows = append(emptyRows, r)
		}
	}

	emptyCols := []int{}
	for i := 0; i < len(grid[0]); i++ {
		col := []byte{}
		for j := 0; j < len(grid); j++ {
			col = append(col, grid[j][i])
		}
		if All(col, func(b byte) bool { return b == '.' }) {
			emptyCols = append(emptyCols, i)
		}
	}

	points := []Coord{}
	for r, row := range grid {
		for c, ch := range row {
			if ch == '#' {
				points = append(points, Coord{R: r, C: c})
			}
		}
	}

	total := 0
	for i, p1 := range points {
		for _, p2 := range points[:i] {
			total += calc(p1.R, p2.R, emptyRows, scale)
			total += calc(p1.C, p2.C, emptyCols, scale)
		}
	}

	return fmt.Sprint(total)
}

func calc(p1, p2 int, empty []int, scale int) int {
	sum := 0
	for p := min(p1, p2); p < max(p1, p2); p++ {
		if slices.Contains(empty, p) {
			sum += scale
		} else {
			sum += 1
		}
	}
	return sum
}

func All[T any](values []T, condition func(T) bool) bool {
	for i := range values {
		if !condition(values[i]) {
			return false
		}
	}
	return true
}

func Lines(input string) []string {
	sc := bufio.NewScanner(strings.NewReader(input))
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}
