package main

import (
	"fmt"
	"sort"
)

func SolveB(input string) string {
	data, S, E := ParseInput(input)
	starts := make([]Coord, 0)
	for y, r := range data {
		for x, rr := range r {
			if rr == 'a' {
				starts = append(starts, Coord{x: x, y: y})
			}
		}
	}
	counts := make([]int, 0)
	for _, start := range starts {
		hm := NewHeightMap(data, S, E)
		c := hm.Count(start)
		if c > -1 {
			counts = append(counts, c)
		}
	}
	sort.Ints(counts)
	return fmt.Sprint(counts[0])
}
