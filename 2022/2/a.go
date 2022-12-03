package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Move int

const (
	_ Move = iota
	Rock
	Paper
	Scissors
)

type Result int

const (
	Lose Result = 0
	Draw Result = 3
	Win  Result = 6
)

func Beats(a, b Move) Result {
	switch a {
	case Rock:
		if b == Paper {
			return Lose
		}
		if b == Scissors {
			return Win
		}
	case Paper:
		if b == Scissors {
			return Lose
		}
		if b == Rock {
			return Win
		}
	case Scissors:
		if b == Rock {
			return Lose
		}
		if b == Paper {
			return Win
		}
	}
	return Draw
}

func strToMove(str string) Move {
	switch str {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	}
	return 0
}

func SolveA(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	score := 0

	for sc.Scan() {
		txt := sc.Text()

		split := strings.Split(txt, " ")

		op, resp := strToMove(split[0]), strToMove(split[1])

		result := Beats(resp, op)

		score += int(result) + int(resp)
	}

	return fmt.Sprint(score)
}
