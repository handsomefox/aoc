package main

import (
	"fmt"
	"strings"

	"aoc/shared"
)

const (
	year = 2025
	day  = 6
)

func main() {
	input := shared.MustReadInput(year, day)

	fmt.Println("Part A:", PartA(input))
	fmt.Println("Part B:", PartB(input))
}

func PartA(input string) string {
	lineCh := shared.LinesBackward(input)
	signs := strings.Fields(<-lineCh)

	total := make([]int, len(signs))
	for line := range lineCh {
		numbers := shared.FieldsInt(line)
		for i, num := range numbers {
			if total[i] == 0 {
				total[i] = num
				continue
			}
			switch signs[i] {
			case "*":
				total[i] *= num
			case "+":
				total[i] += num
			}
		}
	}

	return shared.String(shared.Sum(total))
}

func PartB(input string) string {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	if len(lines) == 0 {
		return "0"
	}

	grid := padGrid(lines)
	opLine := len(grid) - 1
	sum := 0

	currentNums := make([]int, 0)
	currentOp := rune(0)

	for col := len(grid[0]) - 1; col >= 0; col-- {
		if isSeparator(grid, col) {
			sum += apply(currentNums, currentOp)
			currentNums = currentNums[:0]
			currentOp = 0
			continue
		}

		if num := columnNumber(grid, col, opLine); num >= 0 {
			currentNums = append(currentNums, num)
		}

		op := grid[opLine][col]
		if op == '*' || op == '+' {
			currentOp = op
		}
	}

	sum += apply(currentNums, currentOp)

	return shared.String(sum)
}

func padGrid(lines []string) [][]rune {
	maxWidth := 0
	for _, l := range lines {
		maxWidth = max(maxWidth, len(l))
	}

	grid := make([][]rune, len(lines))
	for i, l := range lines {
		if len(l) < maxWidth {
			l += strings.Repeat(" ", maxWidth-len(l))
		}
		grid[i] = []rune(l)
	}
	return grid
}

func isSeparator(grid [][]rune, col int) bool {
	for row := range grid {
		if grid[row][col] != ' ' {
			return false
		}
	}
	return true
}

func columnNumber(grid [][]rune, col, opLine int) int {
	str := strings.Builder{}
	for row := range opLine {
		ch := grid[row][col]
		if ch >= '0' && ch <= '9' {
			str.WriteRune(ch)
		}
	}
	if str.Len() == 0 {
		return -1
	}
	return shared.MustParseInt(str.String())
}

func apply(nums []int, op rune) int {
	if len(nums) == 0 || op == 0 {
		return 0
	}
	acc := nums[0]
	for _, n := range nums[1:] {
		switch op {
		case '*':
			acc *= n
		case '+':
			acc += n
		}
	}
	return acc
}
