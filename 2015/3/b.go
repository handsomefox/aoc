package main

import "fmt"

func SolveB(input string) string {
	visited := make(map[Position]struct{})

	santa := Position{x: 0, y: 0}
	robot := Position{x: 0, y: 0}

	visited[santa] = struct{}{}

	for i, r := range input {
		var modified *Position

		if i%2 == 0 {
			modified = &santa
		} else {
			modified = &robot
		}

		switch r {
		case '^':
			modified.y++
		case 'v':
			modified.y--
		case '>':
			modified.x++
		case '<':
			modified.x--
		}
		visited[*modified] = struct{}{}
	}

	return fmt.Sprint(len(visited))
}
