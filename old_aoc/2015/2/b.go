package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

func Min(nums ...int) (min int) {
	if len(nums) == 0 {
		return 0
	}
	min = math.MaxInt
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func ribbon(l, w, h int) int {
	p1 := 2 * (l + w)
	p2 := 2 * (l + h)
	p3 := 2 * (w + h)
	return Min(p1, p2, p3) + (l * w * h)
}

func SolveB(input string) string {
	var (
		sc    = bufio.NewScanner(strings.NewReader(input))
		total = 0
	)
	for sc.Scan() {
		var (
			text  = sc.Text()
			split = strings.Split(text, "x")
		)
		var (
			l = tryAtoi(split[0])
			w = tryAtoi(split[1])
			h = tryAtoi(split[2])
		)
		total += ribbon(l, w, h)
	}

	return fmt.Sprint(total)
}
