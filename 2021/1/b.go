package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Window struct {
	sum   int
	count int
}

func SolveB(in string) string {
	scanner := bufio.NewScanner(strings.NewReader(in))
	input := make([]int, 0)
	for scanner.Scan() {
		conv, _ := strconv.Atoi(scanner.Text())
		input = append(input, conv)
	}

	windows := make([]Window, 0)

	for i := 0; i < len(input)-2; i++ {
		w := Window{}
		for j := 0; j < 3; j++ {
			w.sum += input[i+j]
			w.count++
		}
		windows = append(windows, w)
	}

	var (
		incr int
		decr int
	)

	for i := 1; i < len(windows); i++ {
		if windows[i-1].sum < windows[i].sum {
			incr++
		} else {
			decr++
		}
	}
	return fmt.Sprintf("Incr: %d, Decr: %d", incr, decr)
}
