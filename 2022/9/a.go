package main

import (
	"fmt"
)

func SolveA(input string) string {
	var (
		moves = MustParseInput(input)
		rope  = NewRope(1)
	)
	for _, move := range moves {
		rope.Move(move)
	}
	return fmt.Sprint(rope.Len())
}
