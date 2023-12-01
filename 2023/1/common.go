package main

import "strconv"

func MustParse(integer string) int {
	i, _ := strconv.Atoi(integer)
	return i
}
