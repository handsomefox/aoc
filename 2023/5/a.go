package main

import (
	"fmt"
	"math"
	"strings"
)

func SolveA(input string) string {
	var (
		split = strings.Split(input, "\n\n")
		seeds = ParseSeeds(split)    // Seeds to plant
		maps  = ParseMaps(split[1:]) // Maps to use
	)

	lowest := math.MaxInt
	for _, s := range seeds {
		lowest = min(SeedToLocation(s, maps), lowest)
	}

	return fmt.Sprint(lowest)
}

func SeedToLocation(seed int, maps map[MapType][]Range) int {
	prev := seed
	for _, t := range MapTypes {
		value := MapValue(prev, maps[t])
		prev = value
	}
	return prev
}

func MapValue(value int, rng []Range) int {
	for _, r := range rng {
		if v, ok := r.SourceToDest(value); ok {
			return v
		}
	}
	return value
}
