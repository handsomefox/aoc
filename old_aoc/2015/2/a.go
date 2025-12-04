package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func smallestSide(l, w, h int) int {
	var s [3]int
	s[0], s[1], s[2] = l*w, w*h, h*l
	sort.IntSlice(s[:]).Sort()
	return s[0]
}

func tryAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func surfaceArea(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

func SolveA(input string) string {
	var (
		sc    = bufio.NewScanner(strings.NewReader(input))
		total = 0
	)
	for sc.Scan() {
		var (
			text  = sc.Text()
			split = strings.Split(text, "x")
		)
		var (
			l = tryAtoi(split[0])
			w = tryAtoi(split[1])
			h = tryAtoi(split[2])
		)
		total += surfaceArea(l, w, h) + smallestSide(l, w, h)
	}

	return fmt.Sprint(total)
}
