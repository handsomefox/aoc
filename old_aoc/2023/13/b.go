package main

import (
	"fmt"
	"slices"
	"strings"
)

func SolveB(input string) string {
	puzzles := strings.Split(input, "\n\n")
	total := 0
	for _, p := range puzzles {
		if p[len(p)-1] == '\n' {
			p = p[:len(p)-1]
		}

		rows := strings.Split(p, "\n")

		row := countB(rows)
		total += 100 * row

		col := countB(parseCols(rows))
		total += col
	}
	return fmt.Sprint(total)
}

func countB(rows []string) int {
	for r := 1; r < len(rows); r++ {
		rev := SliceCopy(rows[:r])
		slices.Reverse(rev)

		top := rev
		bottom := rows[r:]

		length := min(len(top), len(bottom))
		top = top[:length]
		bottom = bottom[:length]

		total := 0
		for i := range top {
			x, y := top[i], bottom[i]
			for i := 0; i < len(x); i++ {
				c1, c2 := x[i], y[i]
				if c1 != c2 {
					total += 1
				}
			}
		}

		if total == 1 {
			return r
		}
	}

	return 0
}
