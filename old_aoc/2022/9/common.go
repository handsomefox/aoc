package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Move struct {
	Direction string
	Amount    int
}

type Point struct {
	x, y int
}

type Rope struct {
	visited map[Point]struct{}
	tail    []Point
	head    Point
}

func NewRope(tails int) *Rope {
	return &Rope{
		visited: make(map[Point]struct{}),
		head:    Point{},
		tail:    make([]Point, tails),
	}
}

func (r Rope) Len() int {
	return len(r.visited)
}

func (r *Rope) Move(m Move) {
	var p Point
	switch m.Direction {
	case "U":
		p = Point{x: 0, y: 1}
	case "D":
		p = Point{x: 0, y: -1}
	case "R":
		p = Point{x: 1, y: 0}
	case "L":
		p = Point{x: -1, y: 0}
	}
	r.move(m.Amount, p)
}

func (r *Rope) move(amount int, p Point) {
	for i := 0; i < amount; i++ {
		r.head.x += p.x
		r.head.y += p.y
		for j := 0; j < len(r.tail); j++ {
			var dx, dy int
			if j == 0 {
				dx, dy = r.head.x-r.tail[0].x, r.head.y-r.tail[0].y
			} else {
				dx, dy = r.tail[j-1].x-r.tail[j].x, r.tail[j-1].y-r.tail[j].y
			}
			if abs(dx) <= 1 && abs(dy) <= 1 {
				continue
			}
			r.tail[j].x += sign(dx)
			r.tail[j].y += sign(dy)
		}
		r.visited[r.tail[len(r.tail)-1]] = struct{}{}
	}
}

func MustParseInput(input string) []Move {
	var (
		moves = make([]Move, 0)
		sc    = bufio.NewScanner(strings.NewReader(input))
	)
	for sc.Scan() {
		split := strings.Split(sc.Text(), " ")
		direction, amount := split[0], MustParseInt(split[1])
		moves = append(moves, Move{Direction: direction, Amount: amount})
	}
	return moves
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func sign(n int) int {
	switch {
	case n == 0:
		return 0
	case n > 0:
		return 1
	default:
		return -1
	}
}
