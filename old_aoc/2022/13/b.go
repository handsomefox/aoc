package main

import (
	"fmt"
	"sort"
)

func SolveB(input string) string {
	pairs := ParseInput(input)

	fmt.Println(pairs)

	sort.Slice(pairs, func(i, j int) bool {
		return Compare(pairs[i], pairs[j]) > 0
	})

	return fmt.Sprint(pairs)
}
