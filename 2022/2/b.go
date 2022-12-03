package main

import (
	"bufio"
	"fmt"
	"strings"
)

func DesiredResult(a Move, r Result) Move {
	switch r {
	case Win:
		if a == Rock {
			return Paper
		}
		if a == Paper {
			return Scissors
		}
		if a == Scissors {
			return Rock
		}
	case Lose:
		if a == Rock {
			return Scissors
		}
		if a == Paper {
			return Rock
		}
		if a == Scissors {
			return Paper
		}
	case Draw:
		return a
	}

	return 0
}

func strToResult(str string) Result {
	switch str {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	}
	return 0
}

func SolveB(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	score := 0
	for sc.Scan() {
		txt := sc.Text()
		split := strings.Split(txt, " ")
		op, desired := strToMove(split[0]), strToResult(split[1])
		ourMove := DesiredResult(op, desired)
		result := Beats(ourMove, op)

		if result == desired {
			score += int(result) + int(ourMove)
		}
	}
	return fmt.Sprint(score)
}
