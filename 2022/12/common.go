package main

import (
	"bufio"
	"strings"
)

func ParseInput(input string) ([][]rune, Coord, Coord) {
	var (
		sc   = bufio.NewScanner(strings.NewReader(input))
		data = make([][]rune, 0)
		S, E Coord
	)
	for i := 0; sc.Scan(); i++ {
		s := make([]rune, 0)
		for j, r := range sc.Text() {
			if string(r) == "S" {
				S = Coord{x: j, y: len(data)}
				s = append(s, 'a')
				continue
			}
			if string(r) == "E" {
				E = Coord{x: j, y: len(data)}
				s = append(s, 'z')
				continue
			}
			s = append(s, r)
		}
		data = append(data, s)
	}
	return data, S, E
}

type Coord struct {
	x, y int
}

func (c Coord) Equals(o Coord) bool {
	return c.x == o.x && c.y == o.y
}

func (c Coord) CanMoveTo(o Coord, hm *HeightMap) bool {
	s := hm.Data[c.y][c.x]
	e := hm.Data[o.y][o.x]
	return s+1 >= e
}

type HeightMap struct {
	Data [][]rune
	S, E Coord

	queue                 []Coord
	moveCount, left, next int
	end                   bool
	visited               map[Coord]struct{}
}

func NewHeightMap(Data [][]rune, S, E Coord) *HeightMap {
	return &HeightMap{
		Data:    Data,
		S:       S,
		E:       E,
		queue:   make([]Coord, 0),
		visited: make(map[Coord]struct{}),
	}
}

func (hm *HeightMap) Count(start Coord) int {
	hm.queue = append(hm.queue, start)
	hm.visited[start] = struct{}{}
	hm.left = 1
	for len(hm.queue) > 0 {
		c := hm.queue[0]
		hm.queue = hm.queue[1:]
		if c.Equals(hm.E) {
			hm.end = true
			break
		}
		for _, adj := range hm.adjacencyList(c) {
			_, didVisit := hm.visited[adj]
			if didVisit {
				continue
			}
			hm.next++
			hm.queue = append(hm.queue, adj)
			hm.visited[adj] = struct{}{}
		}
		hm.left--
		if hm.left == 0 {
			hm.left = hm.next
			hm.next = 0
			hm.moveCount++
		}
	}
	if hm.end {
		return hm.moveCount
	}
	return -1
}

func (hm *HeightMap) adjacencyList(curr Coord) []Coord {
	var (
		w, h     = len(hm.Data[0]), len(hm.Data)
		adjacent = make([]Coord, 0, 4)
	)
	coords := [...]Coord{
		{x: curr.x - 1, y: curr.y},
		{x: curr.x + 1, y: curr.y},
		{x: curr.x, y: curr.y - 1},
		{x: curr.x, y: curr.y + 1},
	}
	if curr.x > 0 && curr.CanMoveTo(coords[0], hm) {
		adjacent = append(adjacent, coords[0])
	}
	if curr.x < w-1 && curr.CanMoveTo(coords[1], hm) {
		adjacent = append(adjacent, coords[1])
	}
	if curr.y > 0 && curr.CanMoveTo(coords[2], hm) {
		adjacent = append(adjacent, coords[2])
	}
	if curr.y < h-1 && curr.CanMoveTo(coords[3], hm) {
		adjacent = append(adjacent, coords[3])
	}
	return adjacent
}
