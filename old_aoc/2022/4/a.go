package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Order struct {
	start, end int
}

func NewOrderFromString(input string) Order {
	nums := strings.Split(input, "-")
	return Order{
		start: MustParseInt(nums[0]),
		end:   MustParseInt(nums[1]),
	}
}

func (o Order) Contains(other Order) bool {
	return o.start <= other.start && o.end >= other.end
}

func MustParseInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

func SolveA(input string) string {
	sc := bufio.NewScanner(strings.NewReader(input))

	counter := 0
	for sc.Scan() {
		pair := strings.Split(sc.Text(), ",")
		var (
			first  = NewOrderFromString(pair[0])
			second = NewOrderFromString(pair[1])
		)
		if first.Contains(second) || second.Contains(first) {
			counter++
		}
	}

	return fmt.Sprint(counter)
}
