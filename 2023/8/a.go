package main

import (
	"fmt"
)

func SolveA(input string) string {
	var (
		direction = ParseDirection(input)
		nodes     = ParseNodes(input)
	)

	moves := 0
	node := nodes["AAA"]
	i := 0
	for node.Identifier != "ZZZ" {
		currDirection := direction[i]
		switch string(currDirection) {
		case "R":
			node = node.Right
		case "L":
			node = node.Left
		}
		i += 1
		i %= len(direction)
		moves++
	}

	return fmt.Sprint(moves)
}
