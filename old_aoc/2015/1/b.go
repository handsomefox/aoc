package main

import "fmt"

func SolveB(input string) string {
	var floor int
	for i, v := range input {
		switch v {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			return fmt.Sprint(i + 1)
		}
	}

	return ""
}
