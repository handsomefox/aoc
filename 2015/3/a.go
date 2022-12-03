package main

import "fmt"

type Position struct {
	x, y int
}

func SolveA(input string) string {
	visited := make(map[Position]struct{})
	cur := Position{x: 0, y: 0}
	visited[cur] = struct{}{}

	for _, r := range input {
		switch r {
		case '^':
			cur.y++
		case 'v':
			cur.y--
		case '>':
			cur.x++
		case '<':
			cur.x--
		}
		visited[cur] = struct{}{}
	}

	return fmt.Sprint(len(visited))
}
