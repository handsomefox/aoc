package main

import (
	"fmt"
)

func SolveA(input string) string {
	pairs := ParseInput(input)
	indices := 0

	for i, pair := range pairs {
		result := Compare(pair.A, pair.B)
		if result == -1 {
			indices += i + 1
		}
	}

	return fmt.Sprint(indices)
}
