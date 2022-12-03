package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Submarine2 struct {
	h, depth, aim int
}

func (s *Submarine2) exec(c Command, val int) {
	switch c {
	case Forward:
		s.h += val
		s.depth += s.aim * val
	case Down:
		s.aim += val
	case Up:
		s.aim -= val
	}
}

func parseCommand2(line string) (Command, int) {
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

func SolveB(input string) string {
	submarine := new(Submarine2)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		cmd, val := parseCommand2(scanner.Text())
		submarine.exec(cmd, val)
	}
	return fmt.Sprint(submarine.depth * submarine.h)
}
