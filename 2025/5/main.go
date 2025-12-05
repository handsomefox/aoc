package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"aoc/shared"
)

const (
	year = 2025
	day  = 5
)

func main() {
	input := shared.MustReadInput(year, day)

	fmt.Println("Part A:", PartA(input))
	fmt.Println("Part B:", PartB(input))
}

func PartA(input string) string {
	puzzleInput := parse(input)
	return shared.String(puzzleInput.availableInRange())
}

func PartB(input string) string {
	puzzleInput := parse(input)
	return shared.String(puzzleInput.totalInRange())
}

type Puzzle struct {
	Fresh     []shared.Tuple[int]
	Available []int
}

func parse(input string) Puzzle {
	split := strings.Split(input, "\n\n")
	fr := strings.Split(split[0], "\n")
	avail := strings.Split(split[1], "\n")

	i := Puzzle{
		Fresh:     make([]shared.Tuple[int], 0, len(fr)),
		Available: shared.LinesAsInts(avail),
	}

	slices.Sort(fr)
	fr = slices.Compact(fr)

	for _, r := range fr {
		split := strings.Split(r, "-")
		start, end := split[0], split[1]
		i.Fresh = append(i.Fresh, shared.NewTuple(shared.MustParseInt(start), shared.MustParseInt(end)))
	}

	return i
}

func (i *Puzzle) availableInRange() int {
	total := 0
	for _, value := range i.Available {
		for _, r := range i.Fresh {
			if value >= r.First && value <= r.Second {
				total++
				break
			}
		}
	}
	return total
}

func (i *Puzzle) totalInRange() int {
	i.merge()
	total := 0
	for _, r := range i.Fresh {
		total += r.Second - r.First + 1
	}
	return total
}

func (i *Puzzle) merge() {
	slices.SortFunc(i.Fresh, func(a, b shared.Tuple[int]) int {
		return cmp.Compare(a.First, b.First)
	})
	merged := make([]shared.Tuple[int], 0, len(i.Fresh))

	for _, r := range i.Fresh {
		if len(merged) == 0 || r.First > merged[len(merged)-1].Second+1 {
			merged = append(merged, r)
		} else if r.Second > merged[len(merged)-1].Second {
			merged[len(merged)-1].Second = r.Second
		}
	}

	i.Fresh = merged
}
