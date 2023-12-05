package main

import (
	"strconv"
	"strings"
)

type Range struct {
	Destination, Source, Length int
}

func (r Range) SourceToDest(src int) (int, bool) {
	if src >= r.Source && src <= r.Source+r.Length-1 {
		diff := src - r.Source
		dest := r.Destination + diff
		if dest >= r.Destination && dest <= r.Destination+r.Length-1 {
			return dest, true
		}
	}
	return 0, false
}

func (r Range) DestToSource(dest int) (int, bool) {
	if dest >= r.Destination && dest <= r.Destination+r.Length-1 {
		diff := dest - r.Destination
		src := r.Source + diff
		if src >= r.Source && src <= r.Source+r.Length-1 {
			return src, true
		}
	}
	return 0, false
}

func NewMap(input string) []Range {
	// First line is guaranteed to be something like soil-to-fertilizer map:
	lines := strings.Split(input, "\n")
	lines = lines[1:] // Remove the line with the map type

	ranges := make([]Range, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		dest := fields[0]
		src := fields[1]
		length := fields[2]
		ranges = append(ranges, Range{
			Destination: MustParse(dest),
			Source:      MustParse(src),
			Length:      MustParse(length),
		})
	}

	return ranges
}

func ParseSeeds(inputs []string) []int {
	seeds := make([]int, 0)
	for _, field := range strings.Fields(strings.Split(inputs[0], "seeds: ")[1]) {
		seeds = append(seeds, MustParse(field))
	}
	return seeds
}

func ParseMaps(inputs []string) [][]Range {
	maps := make([][]Range, 0)
	for _, s := range inputs {
		maps = append(maps, NewMap(s))
	}
	return maps
}

func MustParse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int(i)
}
