package main

import (
	"fmt"
	"math"
	"strings"
)

func SolveA(input string) string {
	var (
		split  = strings.Split(input, "\n\n")
		seeds  = ParseSeeds(split)    // Seeds to plant
		ranges = ParseMaps(split[1:]) // Maps to use
	)

	lowest := math.MaxInt
	for i := 0; i < len(seeds); i++ {
		lowest = min(SeedToLocation(seeds[i], ranges), lowest)
	}

	return fmt.Sprint(lowest)
}

func SeedToLocation(seed int, ranges [][]Range) int {
	value := seed
	for i := 0; i < len(ranges); i++ {
		value = MapValue(value, ranges[i])
	}
	return value
}

func MapValue(value int, rng []Range) int {
	for i := 0; i < len(rng); i++ {
		if v, ok := rng[i].SourceToDest(value); ok {
			return v
		}
	}
	return value
}
