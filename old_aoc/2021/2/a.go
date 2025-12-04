package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Command uint8

const (
	Forward Command = iota
	Down
	Up
)

type Submarine struct {
	h, depth int
}

func (s *Submarine) exec(c Command, val int) {
	switch c {
	case Forward:
		s.h += val
	case Down:
		s.depth += val
	case Up:
		s.depth -= val
	}
}

func parseCommand(line string) (Command, int) {
	split := strings.Fields(line)

	value, _ := strconv.Atoi(split[1])
	var cmd Command

	commandString := split[0]
	switch commandString {
	case "forward":
		cmd = Forward
	case "down":
		cmd = Down
	case "up":
		cmd = Up
	}
	return cmd, value
}

func SolveA(input string) string {
	submarine := new(Submarine)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		cmd, val := parseCommand(scanner.Text())
		submarine.exec(cmd, val)
	}
	return fmt.Sprint(submarine.depth * submarine.h)
}
