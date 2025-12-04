package main

import (
	. "common"
	"strconv"
)

func SolveB(input string) string {
	Assert(len(input) != 0, "input must not be empty")

	tl, tr := []int{}, []int{}
	for line := range Lines(input) {
		l, r := ParseInput(line)
		tl = append(tl, l)
		tr = append(tr, r)
	}

	for i := range tr {
		tl[i] *= Instances(tr, tl[i])
	}

	return strconv.FormatInt(int64(Sum(tl)), 10)
}
