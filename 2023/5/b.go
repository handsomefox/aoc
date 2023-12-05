package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
)

type Seed struct {
	Start, Length int
}

func SolveB(input string) string {
	var (
		split       = strings.Split(input, "\n\n")
		parsedSeeds = ParseSeeds(split)    // Seeds to plant
		ranges      = ParseMaps(split[1:]) // Maps to use
	)
	slices.Reverse(ranges)

	seeds := make([]Seed, 0)
	// it looks like the seeds: line actually describes ranges of seed numbers.
	for i := 0; i < len(parsedSeeds); i += 2 {
		seeds = append(seeds, Seed{
			Start:  parsedSeeds[i],
			Length: parsedSeeds[i+1],
		})
	}

	for location := 0; location < math.MaxInt; location++ {
		if seed, ok := LocationToSeed(location, ranges); ok {
			for _, s := range seeds {
				if seed >= s.Start && seed <= s.Start+s.Length-1 {
					return fmt.Sprint(location)
				}
			}
		}
	}

	return ""
}

func LocationToSeed(location int, ranges [][]Range) (int, bool) {
	value := location
	ok := false
	for i := 0; i < len(ranges); i++ {
		value, ok = MapValueChecked(value, ranges[i])
		if !ok {
			return value, false
		}
	}
	return value, ok
}

func MapValueChecked(value int, rng []Range) (int, bool) {
	for i := 0; i < len(rng); i++ {
		if v, ok := rng[i].DestToSource(value); ok {
			return v, true
		}
	}
	return value, false
}
