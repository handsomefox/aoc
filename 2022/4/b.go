package main

import (
	"bufio"
	"fmt"
	"strings"
)

func (o Order) OverlapsWith(other Order) bool {
	m := make(map[int]int)

	for i := o.start; i <= o.end; i++ {
		m[i]++
	}
	for i := other.start; i <= other.end; i++ {
		m[i]++
	}
	for _, v := range m {
		if v > 1 {
			return true
		}
	}

	return false
}

func SolveB(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	counter := 0
	for sc.Scan() {
		pair := strings.Split(sc.Text(), ",")
		var (
			first  = NewOrderFromString(pair[0])
			second = NewOrderFromString(pair[1])
		)
		if first.OverlapsWith(second) {
			counter++
		}
	}

	return fmt.Sprint(counter)
}
