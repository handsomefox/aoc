package main

import (
	"strconv"
)

func AmountOfWinningNumbers(time, distance int) int {
	nums := make([]int, 0)
	for speed := 1; speed < time/2+1; speed++ {
		travel := time - speed
		totalTraveled := travel * speed
		if totalTraveled > distance {
			nums = append(nums, totalTraveled)
		}
	}
	return len(nums) * 2
}

func MustParse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func IntSlice(s []string) []int {
	ints := make([]int, 0, len(s))
	for i := range s {
		ints = append(ints, MustParse(s[i]))
	}
	return ints
}
