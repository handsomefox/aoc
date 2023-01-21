package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Monkey struct {
	id            int
	startingItems []int
	operation     func(i int) int
	test          int
	ifTrue        int
	ifFalse       int

	inspected int
}

func ParseMonkeys(input string) []*Monkey {
	var (
		sc      = bufio.NewScanner(strings.NewReader(input))
		monkeys = make([]*Monkey, 0)
	)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "Monkey ") {
			m := MakeMonkey(sc)
			monkeys = append(monkeys, m)
		}
	}
	return monkeys
}

func MakeMonkey(sc *bufio.Scanner) *Monkey {
	id := ParseID(sc.Text())
	sc.Scan()

	items := ParseItems(sc.Text())
	sc.Scan()

	op := ParseOP(sc.Text())
	sc.Scan()

	test := ParseTest(sc.Text())
	sc.Scan()

	ifTrue := ParseTrue(sc.Text())
	sc.Scan()

	ifFalse := ParseFalse(sc.Text())

	return &Monkey{
		id:            id,
		startingItems: items,
		operation:     op,
		test:          test,
		ifTrue:        ifTrue,
		ifFalse:       ifFalse,
	}
}

func ParseID(s string) int {
	_, after, _ := strings.Cut(s, "Monkey ")
	before, _, _ := strings.Cut(after, ":")
	return MustParseInt(before)
}

func ParseItems(s string) []int {
	itemsStr := strings.Split(s, "  Starting items: ")[1]
	items := strings.Split(itemsStr, ", ")
	i := make([]int, 0, len(items))
	for _, v := range items {
		i = append(i, MustParseInt(v))
	}
	return i
}

func ParseOP(s string) func(i int) int {
	op := strings.Split(s, "  Operation: new = ")[1]

	const old = "old"

	if strings.Contains(op, "+") {
		v := strings.Split(op, " + ")[1]
		if v == old {
			return func(i int) int {
				return i + i
			}
		}
		return func(i int) int {
			return i + MustParseInt(v)
		}
	}
	if strings.Contains(op, "-") {
		v := strings.Split(op, " - ")[1]
		if v == old {
			return func(i int) int {
				return 0
			}
		}
		return func(i int) int {
			return i - MustParseInt(v)
		}
	}
	if strings.Contains(op, "*") {
		v := strings.Split(op, " * ")[1]
		if v == old {
			return func(i int) int {
				return i * i
			}
		}
		return func(i int) int {
			return i * MustParseInt(v)
		}
	}
	if strings.Contains(op, "/") {
		v := strings.Split(op, " / ")[1]
		if v == old {
			return func(i int) int {
				return 1
			}
		}
		return func(i int) int {
			return i / MustParseInt(v)
		}
	}

	panic("failed to parse op")
}

func ParseTest(s string) int {
	return MustParseInt(strings.Split(s, "  Test: divisible by ")[1])
}

func ParseTrue(s string) int {
	return MustParseInt(strings.Split(s, "    If true: throw to monkey ")[1])
}

func ParseFalse(s string) int {
	return MustParseInt(strings.Split(s, "    If false: throw to monkey ")[1])
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
