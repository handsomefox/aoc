package main

import (
	"fmt"
	"slices"
)

func SolveA(input string) string {
	grid := Bytes(input)
	sr, sc := findStart(grid)

	deque := NewDeque[Coords]()
	deque.PushRight(Coords{Row: sr, Col: sc})

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
		}

		down := Coords{Row: r + 1, Col: c}
		if r < len(grid)-1 && CanGoAndReceive(ch, GoDown, grid[r+1][c], RecvDown) && !inSeen(down) {
			seen = append(seen, down)
			deque.PushRight(down)
		}

		left := Coords{Row: r, Col: c - 1}
		if c > 0 && CanGoAndReceive(ch, GoLeft, grid[r][c-1], RecvLeft) && !inSeen(left) {
			seen = append(seen, left)
			deque.PushRight(left)
		}

		right := Coords{Row: r, Col: c + 1}
		if c < len(grid[r])-1 && CanGoAndReceive(ch, GoRight, grid[r][c+1], RecvRight) && !inSeen(right) {
			seen = append(seen, right)
			deque.PushRight(right)
		}
	}

	return fmt.Sprint(len(seen) / 2)
}
