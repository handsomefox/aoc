package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func SolveB(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	nums := make([]int, 0)
	for sc.Scan() {
		txt := sc.Text()
		nums = append(nums, getNumsB(txt))
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return fmt.Sprint(sum)
}

var valid = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getNumsB(input string) int {
	str := ""
	pos := make(map[int]string)
	for i, r := range input {
		if unicode.IsDigit(r) && r != '0' {
			pos[i] = string(r)
		}
	}

	for k := range valid {
		first, last := strings.Index(input, k), strings.LastIndex(input, k)
		if first != -1 {
			pos[first] = valid[k]
		}
		if last != -1 {
			pos[last] = valid[k]
		}
	}

	arr := make([]int, 0)
	for k := range pos {
		arr = append(arr, k)
	}
	sort.IntSlice(arr).Sort()

	if len(arr) >= 2 {
		str = pos[arr[0]] + pos[arr[len(arr)-1]]
	} else if len(arr) == 1 {
		str = pos[arr[0]] + pos[arr[0]]
	} else {
		return 0
	}

	return MustParse(str)
}
