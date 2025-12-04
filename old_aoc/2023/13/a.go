package main

import (
	"fmt"
	"slices"
	"strings"
)

func SolveA(input string) string {
	puzzles := strings.Split(input, "\n\n")
	total := 0
	for _, p := range puzzles {
		if p[len(p)-1] == '\n' {
			p = p[:len(p)-1]
		}

		rows := strings.Split(p, "\n")

		row := count(rows)
		total += 100 * row

		col := count(parseCols(rows))
		total += col
	}
	return fmt.Sprint(total)
}

func count(rows []string) int {
	for r := 1; r < len(rows); r++ {
		rev := SliceCopy(rows[:r])
		slices.Reverse(rev)

		top := rev
		bottom := rows[r:]

		length := min(len(top), len(bottom))
		top = top[:length]
		bottom = bottom[:length]

		if slices.Equal(top, bottom) {
			return r
		}
	}

	return 0
}
