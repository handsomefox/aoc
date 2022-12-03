package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const DAYS = 256

// each index corresponds to to the amount of fishes of the same generation
// if you have fish with a counter of 1, it will be found at index 1
// every day, we shift all the fishes to the left and handle the ones that are at zero'th index
// specifically
// at the end, we have to count the sum of all fishes
var fishes [9]int64

func passDay2(fishes [9]int64) [9]int64 {
	// handle day 0 specifically
	dayZero := fishes[0]
	fishes[0] = 0
	// we remember the number and reset the value in the array
	// each fish creates a new one, so we move that count to index 8
	// and we also move that count to index 6 as the fishes reset their counter

	// handle the other days
	// every day the fishes move to the left
	for i := 1; i < len(fishes); i++ {
		fishes[i-1] = fishes[i]
	}

	// after moving all the fishes to the left, we use the remember values to
	// put fishes from day zero to the end and the reset fishes to day 6
	fishes[6] += dayZero
	fishes[8] = dayZero

	return fishes
}

func SolveB(input string) string {
	sn := bufio.NewScanner(strings.NewReader(input))

	fishMap := make(map[int]int64)

	for sn.Scan() {
		values := strings.Split(sn.Text(), ",")

		for _, value := range values {
			v, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			fishMap[v] = fishMap[v] + 1
		}
	}

	for i := 0; i < len(fishes); i++ {
		fishes[i] = fishMap[i]
	}

	for i := 0; i < DAYS; i++ {
		fishes = passDay2(fishes)
	}

	var sum int64
	for i := 0; i < len(fishes); i++ {
		sum += fishes[i]
	}

	return fmt.Sprint(sum)
}
