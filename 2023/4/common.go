package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	Actual  []int
	Winning []int
	Number  int
}

func (c Card) String() string {
	return fmt.Sprintf("Card #%d: winning{%v} actual{%v}", c.Number, c.Winning, c.Actual)
}

func (c Card) Points() int {
	contained := 0
	for _, actual := range c.Actual {
		if slices.Contains(c.Winning, actual) {
			contained++
		}
	}
	return int(math.Pow(2, float64(contained-1)))
}

func (c Card) Matches() int {
	contained := 0
	for _, actual := range c.Actual {
		if slices.Contains(c.Winning, actual) {
			contained++
		}
	}
	return contained
}

func parseNumbers(s string) []int {
	nums := make([]int, 0)
	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.TrimSpace(s)

	for _, s := range strings.Split(s, " ") {
		nums = append(nums, MustParse(strings.TrimSpace(s)))
	}

	return nums
}

func MustParse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
