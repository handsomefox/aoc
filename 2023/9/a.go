package main

import (
	"fmt"
)

func extrapolate(h []int, backwards bool) int {
	if All(h, func(i int) bool { return i == 0 }) {
		return 0
	}

	deltas := Map(Zip(h, h[1:]), func(z Zipped[int]) int {
		return z.Second - z.First
	})
	diff := extrapolate(deltas, backwards)

	if backwards {
		return h[0] - diff
	}
	return h[len(h)-1] + diff
}

func SolveA(input string) string {
	histories := Map(Lines(input), func(line string) []int {
		return MustParseIntegerSlice[int](line)
	})

	sum := 0
	for _, h := range histories {
		sum += extrapolate(h, false)
	}

	return fmt.Sprint(sum)
}
