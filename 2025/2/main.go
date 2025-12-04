package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"aoc/shared"
)

const (
	year = 2025
	day  = 2
)

func main() {
	input := shared.MustReadInput(year, day)

	fmt.Println("Part A:", PartA(input))
	fmt.Println("Part B:", PartB(input))
}

func PartA(input string) string {
	return solveWith(input, invalidIDs)
}

func PartB(input string) string {
	return solveWith(input, invalidIDsAnyPattern)
}

func solveWith(input string, invalid func(int, int) []int) string {
	ranges := strings.Split(input, ",")
	ids := make([]int, 0)

	for _, rng := range ranges {
		split := strings.Split(rng, "-")
		l, r := shared.MustParseInt(split[0]), shared.MustParseInt(split[1])
		ids = append(ids, invalid(l, r)...)
	}

	slices.Sort(ids)
	ids = slices.Compact(ids)
	sum := shared.Reduce(ids, 0, func(accumulator, value int) int { return accumulator + value })

	return shared.String(sum)
}

func invalidIDs(l, r int) (out []int) {
	maxDigits := len(shared.String(r))

	for k := 1; k*2 <= maxDigits; k++ {
		base := int(math.Pow10(k))
		mult := base + 1
		minHalf := int(math.Pow10(k - 1))
		maxHalf := base - 1

		startHalf := max(minHalf, shared.CeilDiv(l, mult))
		endHalf := min(maxHalf, r/mult)
		for h := startHalf; h <= endHalf; h++ {
			out = append(out, h*mult)
		}
	}
	return out
}

func invalidIDsAnyPattern(l, r int) []int {
	maxDigits := len(shared.String(r))
	seen := make(map[int]struct{})
	out := make([]int, 0)

	for totalDigits := 2; totalDigits <= maxDigits; totalDigits++ {
		for patternDigits := 1; patternDigits*2 <= totalDigits; patternDigits++ {
			if totalDigits%patternDigits != 0 {
				continue
			}
			base := int(math.Pow10(patternDigits))
			minPattern := int(math.Pow10(patternDigits - 1))
			maxPattern := base - 1

			multiplier := (int(math.Pow10(totalDigits)) - 1) / (base - 1)

			startPattern := max(minPattern, shared.CeilDiv(l, multiplier))
			endPattern := min(maxPattern, r/multiplier)
			for p := startPattern; p <= endPattern; p++ {
				n := p * multiplier
				if _, exists := seen[n]; !exists {
					seen[n] = struct{}{}
					out = append(out, n)
				}
			}
		}
	}

	return out
}
