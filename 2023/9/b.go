package main

import (
	"fmt"
)

func SolveB(input string) string {
	histories := Map(Lines(input), func(line string) []int {
		return MustParseIntegerSlice[int](line)
	})

	sum := 0
	for _, h := range histories {
		sum += extrapolate(h, true)
	}

	return fmt.Sprint(sum)
}
