package main

import (
	"crypto/md5"
	"strconv"
	"strings"
)

func SolveA(input string) string {
	input = strings.TrimSpace(input)

	var result string
	for i := 1; ; i++ {
		i := strconv.Itoa(i)
		hash := md5.Sum([]byte(input + i))
		if hash[0] == 0 && hash[1] == 0 && hash[2] < 0x0F {
			result = i
			break
		}
	}

	return result
}
