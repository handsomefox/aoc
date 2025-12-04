package main

import (
	"fmt"
	"strings"
)

func SolveA(input string) string {
	up, down := strings.Count(input, "("), strings.Count(input, ")")
	return fmt.Sprint(up - down)
}
