package main

import (
	"fmt"
	"sort"
)

func GetAllSubstrs(input string, length int, outCh chan<- string, quitCh <-chan struct{}) {
	for i := 0; i < len(input)-length; i++ {
		select {
		case <-quitCh:
			return
		default:
			substr := input[i : i+length]
			if len(substr) < length {
				break
			}
			outCh <- substr
		}
	}
}

func FindFirstUnique(inCh <-chan string) int {
	i := 0
	for substr := range inCh {
		var (
			found = false
			s     = []rune(substr)
		)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		for i := 1; i < len(s); i++ {
			if s[i] == s[i-1] {
				found = false
				break
			} else {
				found = true
			}
		}
		if found {
			return i + len(substr)
		}
		i++
	}
	return 0
}

func Solve(input string, length int) string {
	c := make(chan string)
	quitCh := make(chan struct{})

	go GetAllSubstrs(input, length, c, quitCh)

	i := FindFirstUnique(c)
	return fmt.Sprint(i)
}
