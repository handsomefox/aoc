package main

import (
	"strconv"
	"strings"
)

type MapType string

const (
	SeedToSoil            MapType = "seed-to-soil"
	SoilToFertilizer      MapType = "soil-to-fertilizer"
	FertilizerToWater     MapType = "fertilizer-to-water"
	WaterToLight          MapType = "water-to-light"
	LightToTemperature    MapType = "light-to-temperature"
	TemperatureToHumidity MapType = "temperature-to-humidity"
	HumidityToLocation    MapType = "humidity-to-location"
)

var MapTypes []MapType = []MapType{
	SeedToSoil,
	SoilToFertilizer,
	FertilizerToWater,
	WaterToLight,
	LightToTemperature,
	TemperatureToHumidity,
	HumidityToLocation,
}

type Map struct {
	Type   MapType
	Ranges []Range
}

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

func NewMap(input string) Map {
	// First line is guaranteed to be something like soil-to-fertilizer map:
	lines := strings.Split(input, "\n")
	mapType, _ := strings.CutSuffix(lines[0], " map:")
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

	return Map{
		Type:   MapType(mapType),
		Ranges: ranges,
	}
}

func ParseSeeds(inputs []string) []int {
	seeds := make([]int, 0)
	for _, field := range strings.Fields(strings.Split(inputs[0], "seeds: ")[1]) {
		seeds = append(seeds, MustParse(field))
	}
	return seeds
}

func ParseMaps(inputs []string) map[MapType][]Range {
	maps := make(map[MapType][]Range)
	for _, s := range inputs {
		m := NewMap(s)
		maps[m.Type] = m.Ranges
	}
	return maps
}

func MustParse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
