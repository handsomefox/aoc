package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolveA(in string) string {
	scanner := bufio.NewScanner(strings.NewReader(in))

	input := make([]int, 0)

	for scanner.Scan() {
		conv, _ := strconv.Atoi(scanner.Text())
		input = append(input, conv)
	}

	var (
		bigger  int
		smaller int
	)

	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			bigger++
		} else {
			smaller++
		}
	}
	return fmt.Sprintf("Bigger: %d, Smaller: %d", bigger, smaller)
}
