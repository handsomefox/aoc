package main

import (
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

type Move struct {
	amount int
	from   int
	to     int
}

func Solve(input string, shouldReverse bool) string {
	var (
		sc     = bufio.NewScanner(strings.NewReader(input))
		stacks = make([]string, 0)
		moves  = make([]Move, 0)
	)
	parsedCrates := false
	for sc.Scan() {
		text := sc.Text()
		// parse the initial crates input
		if !parsedCrates {
			if text != "" {
				stacks = append(stacks, text)
			} else {
				parsedCrates = true
			}
			continue
		}
		// Get the moves
		split := strings.Split(text, " ")
		// append the move
		moves = append(moves, Move{
			amount: MustParseInt(split[1]),
			from:   MustParseInt(split[3]),
			to:     MustParseInt(split[5]),
		})
	}
	// make a map for storing the crates
	crateMap := make(map[int][]string)
	nums := stacks[len(stacks)-1]
	for _, v := range nums { // the input is " 1   2   3"....
		if unicode.IsNumber(v) { // ignore whitespace
			i := MustParseInt(string(v))
			crateMap[i] = make([]string, 0)
		}
	}
	// parse the crates to a map
	runeIndex := 1
	for i := 0; i < len(crateMap); i++ { // for each stack
		for j := len(stacks) - 2; j > -1; j-- { // go through the crates, bottom to top
			str := stacks[j] // get the current line
			if runeIndex > len(str) || str[runeIndex] == ' ' {
				continue
			}
			crateMap[i+1] = append(crateMap[i+1], string(str[runeIndex])) // append the expected crate value
		}
		runeIndex += 4 // crates indicies have 3 whitespace characters between them.
	}
	for k := range crateMap {
		crateMap[k] = reverse(crateMap[k]) // reversing them makes performing moves easier.
	}

	crateMap = PerformMoves(moves, crateMap, shouldReverse)

	result := ""
	for i := 0; i < len(crateMap); i++ {
		result += crateMap[i+1][0]
	}
	return result
}

func PerformMoves(moves []Move, crates map[int][]string, shouldReverse bool) map[int][]string {
	for _, move := range moves {
		// Copy the slice from the origin crate.
		from := make([]string, 0)
		from = append(from, crates[move.from]...)
		// Take only the required crates.
		sliced := from[:move.amount]
		// Take the required crates from origin, too.
		crates[move.from] = crates[move.from][move.amount:]
		// Append the destination crates to the sliced crates,
		// P.S. sliced crates are added in reverse order
		// for the first part of the task.
		if shouldReverse {
			sliced = append(reverse(sliced), crates[move.to]...)
		} else {
			sliced = append(sliced, crates[move.to]...)
		}
		// Store the result in the destination
		crates[move.to] = sliced
	}
	return crates
}

func MustParseInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

func reverse[T any](s []T) []T {
	a := make([]T, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
