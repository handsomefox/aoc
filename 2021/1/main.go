package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println("Part 1: ", SolveA(input))
	fmt.Println("Part 2: ", SolveB(input))
}
