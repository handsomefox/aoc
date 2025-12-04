package main

import (
	"fmt"
	"sort"
)

func SolveA(input string) string {
	monkeys := ParseMonkeys(input)

	for round := 0; round < 20; round++ {
		for _, monkey := range monkeys {
			for i := 0; i < len(monkey.startingItems); i++ {
				monkey.inspected++
				new := monkey.operation(monkey.startingItems[i]) / 3
				if new%monkey.test == 0 {
					monkeys[monkey.ifTrue].startingItems = append(monkeys[monkey.ifTrue].startingItems, new)
				} else {
					monkeys[monkey.ifFalse].startingItems = append(monkeys[monkey.ifFalse].startingItems, new)
				}
			}
			monkey.startingItems = monkey.startingItems[:0]
		}
	}

	inspected := make([]int, 0)
	for _, monkey := range monkeys {
		inspected = append(inspected, monkey.inspected)
	}
	sort.Ints(inspected)
	inspected = inspected[len(inspected)-2:]

	return fmt.Sprint(inspected[0] * inspected[1])
}
