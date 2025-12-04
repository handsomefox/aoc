package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SolveA(input string) string {
	r := strings.NewReader(input)
	sc := bufio.NewScanner(r)

	max := math.MinInt
	currentElf := 0
	for sc.Scan() {
		text := sc.Text()

		if text == "" {
			if currentElf > max {
				max = currentElf
			}
			currentElf = 0
			continue
		}

		value, _ := strconv.Atoi(text)
		currentElf += value
	}

	return fmt.Sprint(max)
}
