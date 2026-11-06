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
	day  = 8
)

type Point struct {
	X, Y, Z int
}

type Distance struct {
	A, B  Point
	Value float64
}

func (d Distance) String() string {
	return fmt.Sprintf("{a=%v->b=%v == %v}", d.A, d.B, d.Value)
}

func NewPoint(line string) Point {
	s := strings.Split(line, ",")
	return Point{
		X: shared.MustParseInt(s[0]),
		Y: shared.MustParseInt(s[1]),
		Z: shared.MustParseInt(s[2]),
	}
}

func (p Point) String() string { return fmt.Sprintf("{x:%v y:%v z:%v}", p.X, p.Y, p.Z) }
func (p Point) Equal(other Point) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}

func FindDistance(a, b Point) float64 {
	return shared.Sqrt(shared.Pow((a.X-b.X), 2) +
		shared.Pow((a.Y-b.Y), 2) +
		shared.Pow((a.Z-b.Z), 2),
	)
}

func main() {
	input := shared.MustReadInput(year, day)

	fmt.Println("Part A:", PartA(input))
	fmt.Println("Part B:", PartB(input))
}

func PartA(input string) string {
	// TODO: implement Part A solution

	points := make([]Point, 0)
	distances := make([]Distance, 0)

	for line := range shared.Lines(input) {
		point := NewPoint(line)
		points = append(points, point)
	}

	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			dist := FindDistance(a, b)
			distances = append(distances, Distance{A: a, B: b, Value: dist})
		}
	}

	slices.SortFunc(distances, func(a, b Distance) int {
		return cmp.Compare(a.Value, b.Value)
	})

	circuits := make([][]Point, 0)

	for _, dist := range distances {
		if len(circuits) == 0 {
			circuits = append(circuits, []Point{dist.A, dist.B})
			continue
		}

		for i, circuit := range circuits {
			for _, point := range circuit {
				switch {
				case point.Equal(dist.A):
					circuit = append(circuit, dist.A)
				case point.Equal(dist.B):
					circuit = append(circuit, dist.B)
				}
			}
			circuits[i] = circuit
		}
	}

	return ""
}

func PartB(input string) string {
	// TODO: implement Part B solution
	return ""
}
