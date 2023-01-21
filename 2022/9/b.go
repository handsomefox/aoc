package main

import "fmt"

func SolveB(input string) string {
	var (
		moves = MustParseInput(input)
		rope  = NewRope(9)
	)
	for _, move := range moves {
		rope.Move(move)
	}
	return fmt.Sprint(rope.Len())
}
