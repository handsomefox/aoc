package main

import (
	"fmt"
)

func SolveA(input string) string {
	data, S, E := ParseInput(input)
	hm := NewHeightMap(data, S, E)
	return fmt.Sprint(hm.Count(hm.S))
}
