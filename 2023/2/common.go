package main

import (
	"strconv"
	"strings"
)

// The Elf would first like to know which games would have been possible if the bag
// contained only 12 red cubes, 13 green cubes, and 14 blue cubes?
const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

type (
	Game struct {
		ID       int
		Rounds   []Round
		Possible bool
		Power    int
	}
	Round struct {
		Red, Green, Blue uint8
	}
)

func calculateGame(input string) Game {
	str, _ := strings.CutPrefix(input, "Game ") // Remove game prefix
	split := strings.Split(str, ": ")
	// [0] == ID of the game, [1] == Game rounds
	id := MustParse(split[0])
	rounds := strings.Split(split[1], "; ") // Every round is split up by ;

	game := Game{
		ID:       id,
		Rounds:   make([]Round, 0),
		Possible: false,
	}
	// fmt.Printf("GameID: %d\nRounds: %#v\n", id, rounds)
	for i := 0; i < len(rounds); i++ {
		round := parseRounds(rounds[i])
		game.Rounds = append(game.Rounds, round)
	}
	game.Power = gamePower(game.Rounds)
	game.Possible = isPossible(game.Rounds)
	// fmt.Printf("Game %d: %d\n", game.ID, game.Power)
	return game
}

func parseRounds(input string) Round {
	cubes := strings.Split(input, ", ") // Every cube is split by ,
	// fmt.Printf("Cubes: %#v\n", cubes)
	r := Round{}
	for i := 0; i < len(cubes); i++ {
		redCube := strings.Split(cubes[i], " red")
		greenCube := strings.Split(cubes[i], " green")
		blueCube := strings.Split(cubes[i], " blue")

		if len(redCube) == 2 {
			r.Red = uint8(MustParse(redCube[0]))
		}
		if len(greenCube) == 2 {
			r.Green = uint8(MustParse(greenCube[0]))
		}
		if len(blueCube) == 2 {
			r.Blue = uint8(MustParse(blueCube[0]))
		}
	}
	// fmt.Printf("Round: %#v\n", r)
	return r
}

func gamePower(rounds []Round) int {
	var minRed, minGreen, minBlue uint8
	for _, r := range rounds {
		minRed = max(r.Red, minRed)
		minGreen = max(r.Green, minGreen)
		minBlue = max(r.Blue, minBlue)
	}
	return int(minRed) * int(minGreen) * int(minBlue)
}

func isPossible(rounds []Round) bool {
	for _, r := range rounds {
		if r.Red > MaxRed || r.Green > MaxGreen || r.Blue > MaxBlue {
			return false
		}
	}
	return true
}

func MustParse(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
