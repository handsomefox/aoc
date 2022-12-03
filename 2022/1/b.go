package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SolveB(input string) string {
	r := strings.NewReader(input)
	sc := bufio.NewScanner(r)

	elves := make([]int, 0)

	currentElf := 0
	for sc.Scan() {
		text := sc.Text()

		if text == "" {
			elves = append(elves, currentElf)
			currentElf = 0
			continue
		}

		value, _ := strconv.Atoi(text)
		currentElf += value
	}

	sort.IntSlice(elves).Sort()

	sum := 0
	for _, v := range elves[len(elves)-3:] {
		sum += v
	}

	return fmt.Sprint(sum)
}
