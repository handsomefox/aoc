package main

import "strconv"

func WinningNums(time, distance int) []int {
	nums := make([]int, 0)
	for speed := 1; speed < time; speed++ {
		travel := time - speed
		totalTraveled := travel * speed
		if totalTraveled > distance {
			nums = append(nums, totalTraveled)
		}
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
