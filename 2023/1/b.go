package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveB(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	sum := 0
	for sc.Scan() {
		sum += getNumsB(sc.Text())
	}

	return fmt.Sprint(sum)
}

var numberLookupTable = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

func getNumsB(input string) int {
	var (
		firstIndex = len(input)
		lastIndex  = -1
		firstValue = ""
		lastValue  = ""
	)

	for k, v := range numberLookupTable {
		first, last := strings.Index(input, k), strings.LastIndex(input, k)
		if first != -1 && first < firstIndex {
			firstIndex = first
			firstValue = v
		}
		if last != -1 && last > lastIndex {
			lastIndex = last
			lastValue = v
		}
	}

	return MustParse(firstValue + lastValue)
}
