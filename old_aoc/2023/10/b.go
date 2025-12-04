package main

import (
	"bytes"
	"fmt"
	"slices"
)

func SolveB(input string) string {
	grid := Bytes(input)
	sr, sc := findStart(grid)

	deque := NewDeque[Coords]()
	deque.PushRight(Coords{Row: sr, Col: sc})

	possible := []byte{'|', '-', 'J', 'L', '7', 'F'}

	seen := make([]Coords, 0)
	seen = append(seen, Coords{Row: sr, Col: sc})

	for !deque.Empty() {
		coords := deque.PopLeft()
		r, c := coords.Row, coords.Col
		ch := grid[coords.Row][coords.Col]

		inSeen := func(c Coords) bool { return slices.Contains(seen, c) }

		up := Coords{Row: r - 1, Col: c}
		if r > 0 && CanGoAndReceive(ch, GoUp, grid[r-1][c], RecvUp) && !inSeen(up) {
			seen = append(seen, up)
			deque.PushRight(up)
			if ch == 'S' {
				possible = intersection(possible, GoUp[1:])
			}
		}

		down := Coords{Row: r + 1, Col: c}
		if r < len(grid)-1 && CanGoAndReceive(ch, GoDown, grid[r+1][c], RecvDown) && !inSeen(down) {
			seen = append(seen, down)
			deque.PushRight(down)
			if ch == 'S' {
				possible = intersection(possible, GoDown[1:])
			}
		}

		left := Coords{Row: r, Col: c - 1}
		if c > 0 && CanGoAndReceive(ch, GoLeft, grid[r][c-1], RecvLeft) && !inSeen(left) {
			seen = append(seen, left)
			deque.PushRight(left)
			if ch == 'S' {
				possible = intersection(possible, GoLeft[1:])
			}
		}

		right := Coords{Row: r, Col: c + 1}
		if c < len(grid[r])-1 && CanGoAndReceive(ch, GoRight, grid[r][c+1], RecvRight) && !inSeen(right) {
			seen = append(seen, right)
			deque.PushRight(right)
			if ch == 'S' {
				possible = intersection(possible, GoRight[1:])
			}
		}
	}

	for i := range grid {
		grid[i] = bytes.ReplaceAll(grid[i], []byte{'S'}, []byte{possible[0]})
	}
	inside := make(map[Coords]bool)
	for i := range seen {
		inside[seen[i]] = true
	}
	for r, row := range grid {
		for c := range row {
			if !inside[Coords{Row: r, Col: c}] {
				grid[r][c] = '.'
			}
		}
	}

	outside := make(map[Coords]bool)
	for r, row := range grid {
		var up, within bool
		for c, ch := range row {
			switch {
			case ch == '|':
				within = !within
			case slices.Contains([]byte{'L', 'F'}, ch):
				up = ch == 'L'
			case slices.Contains([]byte{'7', 'J'}, ch):
				var other byte = 'J'
				if !up {
					other = '7'
				}
				if ch != other {
					within = !within
				}
			}
			if !within {
				outside[Coords{Row: r, Col: c}] = true
			}
		}
	}

	return fmt.Sprint(len(grid)*len(grid[0]) - len(union(outside, seen)))
}

func intersection[T comparable](first, second []T) []T {
	out := []T{}
	bucket := map[T]bool{}
	for _, i := range first {
		for _, j := range second {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}
	return out
}

func union(outside map[Coords]bool, seen []Coords) []Coords {
	unionMap := make(map[Coords]bool)
	for elem := range outside {
		unionMap[elem] = true
	}
	for _, elem := range seen {
		unionMap[elem] = true
	}
	var result []Coords
	for key := range unionMap {
		result = append(result, key)
	}
	return result
}
