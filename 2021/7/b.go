package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

func moveCost2(from, to int) int {
	moves := int(math.Abs(float64(from - to)))

	cost := 0
	for i := 1; i <= moves; i++ {
		cost += i
	}

	return cost
}

func SolveB(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	crabPositions := make([]int, 0)
	for scanner.Scan() {
		out := parseInput(scanner.Text())
		crabPositions = append(crabPositions, out...)
	}

	return fmt.Sprintf("minimal move cost would be %v", solve2(crabPositions))
}

func solve2(crabPositions []int) int {
	min, max := MinMax(crabPositions...)

	minCost := math.MaxInt
	for i := min; i < max; i++ {
		totalCost := 0
		for _, pos := range crabPositions {
			totalCost += moveCost2(pos, i)
		}

		if totalCost < minCost {
			minCost = totalCost
		}
	}

	return minCost
}
