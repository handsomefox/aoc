package main

import (
	"crypto/md5"
	"strconv"
	"strings"
)

func SolveB(input string) string {
	input = strings.TrimSpace(input)

	var result string
	for i := 1; ; i++ {
		i := strconv.Itoa(i)
		hash := md5.Sum([]byte(input + i))
		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			result = i
			break
		}
	}

	return result
}
