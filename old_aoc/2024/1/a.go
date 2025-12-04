package main

import (
	. "common"
	"slices"
	"strconv"
	"strings"
)

func ParseInput(line string) (left int, right int) {
	f := strings.Fields(line)
	Assert(len(f) == 2, "there must be only two fields")
	l, r := f[0], f[1]
	return Int[int](l), Int[int](r)
}

func SolveA(input string) string {
	Assert(len(input) != 0, "input must not be empty")

	tl, tr := []int{}, []int{}
	for line := range Lines(input) {
		l, r := ParseInput(line)
		tl = append(tl, l)
		tr = append(tr, r)
	}

	slices.Sort(tl)
	slices.Sort(tr)

	sum := 0
	for _, v := range Zip(tl, tr) {
		sum += Abs(v.First - v.Second)
	}

	return strconv.FormatInt(int64(sum), 10)
}
