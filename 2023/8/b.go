package main

import (
	"fmt"
	"strings"
)

func SolveB(input string) string {
	var (
		direction = ParseDirection(input)
		nodes     = ParseNodes(input)
	)

	startingNodes := make([]*Node, 0)
	for _, node := range nodes {
		if strings.HasSuffix(node.Identifier, "A") {
			startingNodes = append(startingNodes, node)
		}
	}

	return fmt.Sprint(Walk(startingNodes, direction))
}

func Walk(startingNodes []*Node, direction string) int {
	cycles := make([]int, 0)
	for _, node := range startingNodes {
		cycle := make([]int, 0)
		stepCount := 0
		i := 0
		var firstZ *Node
		for {
			for stepCount == 0 || !strings.HasSuffix(node.Identifier, "Z") {
				stepCount += 1
				switch string(direction[i]) {
				case "R":
					node = node.Right
				case "L":
					node = node.Left
				}
				i++
				i %= len(direction)
			}

			cycle = append(cycle, stepCount)

			if firstZ == nil {
				firstZ = node
				stepCount = 0
			} else if node == firstZ {
				break
			}
		}
		cycles = append(cycles, cycle...)
	}

	curr := cycles[0]
	cycles = cycles[2:]
	for i, num := range cycles {
		if i%2 == 0 {
			continue
		}
		curr = (curr * num) / gcd(curr, num)
	}

	return curr
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}
