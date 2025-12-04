package main

import (
	"fmt"
	"strings"
)

func SolveB(input string) string {
	lines := Lines(input)

	total := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		arrangement := ""
		nums := []int{}
		for i := 0; i < 5; i++ {
			arrangement += split[0]
			if i < 4 {
				arrangement += "?"
			}
			nums = append(nums, MustParseSlice(split[1])...)
		}
		total += count(arrangement, nums)
	}

	return fmt.Sprint(total)
}
