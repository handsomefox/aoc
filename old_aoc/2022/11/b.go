package main

import (
	"fmt"
	"sort"
)

func SolveB(input string) string {
	monkeys := ParseMonkeys(input)

	multiple := 1
	for _, m := range monkeys {
		multiple *= m.test
	}

	for round := 0; round < 10000; round++ {
		for _, monkey := range monkeys {
			for i := 0; i < len(monkey.startingItems); i++ {
				monkey.inspected++
				n := monkey.operation(monkey.startingItems[i])
				n %= multiple
				if n%monkey.test == 0 {
					monkeys[monkey.ifTrue].startingItems = append(monkeys[monkey.ifTrue].startingItems, n)
				} else {
					monkeys[monkey.ifFalse].startingItems = append(monkeys[monkey.ifFalse].startingItems, n)
				}
			}
			monkey.startingItems = monkey.startingItems[:0]
		}
	}

	inspected := make([]int, 0, len(monkeys))
	for _, monkey := range monkeys {
		inspected = append(inspected, monkey.inspected)
	}
	sort.Ints(inspected)
	inspected = inspected[len(inspected)-2:]

	return fmt.Sprint(inspected[0] * inspected[1])
}
