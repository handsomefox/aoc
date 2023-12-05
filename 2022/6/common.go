package main

import (
	"fmt"
	"sort"
)

func Solve(input string, length int) string {
	for i := 0; i < len(input)-length; i++ {
		substr := input[i : i+length]
		if len(substr) < length {
			break
		}
		if IsUnique(substr) {
			return fmt.Sprint(i + len(substr))
		}
	}
	return "Failed to solve"
}

func IsUnique(input string) bool {
	s := []rune(input)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return false
		}
	}
	return true
}
