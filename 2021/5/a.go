package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("x: %v, y: %v", p.x, p.y)
}

type Line struct {
	orig, dest Point
}

func (l Line) CoveredPoints() []Point {
	points := make([]Point, 0)

	if l.orig.x == l.dest.x {
		for {
			if l.orig.y > l.dest.y {
				points = append(points, l.dest)
				l.dest.y++
			} else if l.orig.y == l.dest.y {
				points = append(points, l.orig)
				break
			} else {
				points = append(points, l.orig)
				l.orig.y++
			}
		}
	} else if l.orig.y == l.dest.y {
		for {
			if l.orig.x > l.dest.x {
				points = append(points, l.dest)
				l.dest.x++
			} else if l.orig.x == l.dest.x {
				points = append(points, l.orig)
				break
			} else {
				points = append(points, l.orig)
				l.orig.x++
			}
		}
	}
	return points
}

func SolveA(input string) string {
	lines := tryParseInput(strings.NewReader(input))
	coveredPointsCount := make(map[Point]int, 0)
	for _, line := range lines {
		for _, point := range line.CoveredPoints() {
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

func tryParseInteger(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("error parsing integer: %v", err))
	}
	return i
}

func tryParseInput(r io.Reader) []Line {
	var (
		sn    = bufio.NewScanner(r)
		lines = make([]Line, 0)
	)

	// input format: x1,y1 -> x2,y2
	for sn.Scan() {
		line := strings.Split(sn.Text(), " -> ") // Splits into [0](x,y), [1](x,y)
		if len(line) != 2 {
			fmt.Println("Unexpected amount of lines")
			continue
		}

		originSplit := strings.Split(line[0], ",")
		if len(originSplit) != 2 {
			panic("Unexpected amount of points")
		}

		destSplit := strings.Split(line[1], ",")
		if len(destSplit) != 2 {
			panic("Unexpected amount of points")
		}

		x1 := tryParseInteger(originSplit[0])
		y1 := tryParseInteger(originSplit[1])
		origin := Point{x: x1, y: y1}

		x2 := tryParseInteger(destSplit[0])
		y2 := tryParseInteger(destSplit[1])
		dest := Point{x: x2, y: y2}

		lines = append(lines, Line{
			orig: origin,
			dest: dest,
		})
	}
	return lines
}
