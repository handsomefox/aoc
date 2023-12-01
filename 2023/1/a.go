package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

func SolveA(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	nums := make([]int, 0)
	for sc.Scan() {
		txt := sc.Text()
		nums = append(nums, getNumsA(txt))
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return fmt.Sprint(sum)
}

func getNumsA(input string) int {
	var str, first, last string
	for _, r := range input {
		if unicode.IsDigit(r) {
			if first == "" {
				first = string(r)
			}
			last = string(r)
		}
	}

	str = first + last
	return MustParse(str)
}
