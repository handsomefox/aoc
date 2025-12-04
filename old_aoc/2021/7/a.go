package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MinMax(nums ...int) (min int, max int) {
	if len(nums) == 0 {
		return 0, 0
	}

	min = math.MaxInt

	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return
}

func moveCost(from, to int) int {
	return int(math.Abs(float64(from - to)))
}

func SolveA(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	crabPositions := make([]int, 0)
	for scanner.Scan() {
		out := parseInput(scanner.Text())
		crabPositions = append(crabPositions, out...)
	}

	return fmt.Sprintf("minimal move cost would be %v", solve(crabPositions))
}

func parseInput(input string) []int {
	split := strings.Split(input, ",")

	positions := make([]int, 0, len(split))

	for i := 0; i < len(split); i++ {
		integer, err := strconv.Atoi(split[i])
		if err != nil {
			panic(err)
		}

		positions = append(positions, integer)
	}

	return positions
}

func solve(crabPositions []int) int {
	min, max := MinMax(crabPositions...)

	minCost := math.MaxInt
	for i := min; i < max; i++ {
		totalCost := 0
		for _, pos := range crabPositions {
			totalCost += moveCost(pos, i)
		}

		if totalCost < minCost {
			minCost = totalCost
		}
	}

	return minCost
}
