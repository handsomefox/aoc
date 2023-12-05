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
		split = strings.Split(input, "\n\n")
		seeds = ParseSeeds(split)    // Seeds to plant
		maps  = ParseMaps(split[1:]) // Maps to use
	)

	rangeSeeds := make([]Seed, 0)

	// it looks like the seeds: line actually describes ranges of seed numbers.
	for i := 0; i < len(seeds); i += 2 {
		rangeSeeds = append(rangeSeeds, Seed{
			Start:  seeds[i],
			Length: seeds[i+1],
		})
	}

	// Go in reverse (try to find the lowest location number where the seed for it exists)
	mapTypesReverse := make([]MapType, 0)
	mapTypesReverse = append(mapTypesReverse, MapTypes...)
	slices.Reverse(mapTypesReverse)

	for location := 0; location < math.MaxInt; location++ {
		if seed, ok := LocationToSeed(location, maps, mapTypesReverse); ok {
			for _, s := range rangeSeeds {
				if seed >= s.Start && seed <= s.Start+s.Length-1 {
					return fmt.Sprint(location)
				}
			}
		}
	}

	return ""
}

func LocationToSeed(location int, maps map[MapType][]Range, mapTypes []MapType) (int, bool) {
	value := location
	ok := false
	for _, t := range mapTypes {
		value, ok = MapValueChecked(value, maps[t])
	}
	return value, ok
}

func MapValueChecked(value int, rng []Range) (int, bool) {
	for _, r := range rng {
		if v, ok := r.DestToSource(value); ok {
			return v, true
		}
	}
	return value, false
}
