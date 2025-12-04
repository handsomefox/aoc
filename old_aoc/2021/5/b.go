package main

import (
	"fmt"
	"strings"
)

func (p Point) isEqual(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func (line Line) CoveredPoints2() []Point {
	points := make([]Point, 0)
	points = append(points, line.countCoveredPoints()...)
	points = append(points, line.countCoveredDiagonals()...)
	return points
}

func (line Line) countCoveredPoints() []Point {
	points := make([]Point, 0)

	if line.orig.x == line.dest.x {
		for {
			if line.orig.y > line.dest.y {
				points = append(points, line.dest)
				line.dest.y++
			} else if line.orig.y == line.dest.y {
				points = append(points, line.orig)
				break
			} else {
				points = append(points, line.orig)
				line.orig.y++
			}
		}
	} else if line.orig.y == line.dest.y {
		for {
			if line.orig.x > line.dest.x {
				points = append(points, line.dest)
				line.dest.x++
			} else if line.orig.x == line.dest.x {
				points = append(points, line.orig)
				break
			} else {
				points = append(points, line.orig)
				line.orig.x++
			}
		}
	}

	return points
}

func (line Line) countCoveredDiagonals() []Point {
	points := make([]Point, 0)

	// An entry like 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3.
	// An entry like 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.

	s, e := line.orig, line.dest

	if s.x == e.x || s.y == e.y {
		return points
	}

	points = append(points, line.orig)

	for !s.isEqual(e) {
		if s.x > e.x {
			s.x--
		} else if s.x < e.x {
			s.x++
		}

		if s.y > e.y {
			s.y--
		} else if s.y < e.y {
			s.y++
		}

		points = append(points, s)
	}

	return points
}

func SolveB(input string) string {
	lines := tryParseInput(strings.NewReader(input))
	coveredPointsCount := make(map[Point]int, 0)
	for _, line := range lines {
		for _, point := range line.CoveredPoints2() {
			coveredPointsCount[point] = coveredPointsCount[point] + 1
		}
	}

	total := 0
	for _, v := range coveredPointsCount {
		if v > 1 {
			total++
		}
	}
	return fmt.Sprint(total)
}
