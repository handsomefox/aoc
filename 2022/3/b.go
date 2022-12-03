package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveB(input string) string {
	var (
		sc    = bufio.NewScanner(strings.NewReader(input))
		group = make([]string, 0)
		score = 0
	)
	for sc.Scan() {
		txt := sc.Text()
		group = append(group, txt)
		if len(group) != 3 {
			continue
		}

		items := make(map[rune]int)
		for i, v := range group {
			for _, vv := range v {
				value := items[vv]
				if value == i {
					items[vv]++
				}
			}
		}

		for k, v := range items {
			if v == 3 {
				score += GetPriority(k)
			}
		}
		group = make([]string, 0)
	}

	return fmt.Sprint(score)
}
