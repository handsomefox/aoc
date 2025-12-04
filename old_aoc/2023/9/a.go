package main

import (
	"fmt"
)

func extrapolate(slice []int, backwards bool) int {
	if All(slice, func(i int) bool { return i == 0 }) {
		return 0
	}

	deltas := Map(Zip(slice, slice[1:]), func(z Zipped[int]) int {
		return z.Second - z.First
	})
	diff := extrapolate(deltas, backwards)

	if backwards {
		return slice[0] - diff
	}
	return slice[len(slice)-1] + diff
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
