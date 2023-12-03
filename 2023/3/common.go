package main

import (
	"strconv"
	"unicode"
)

type (
	Point struct {
		X, Y int
	}
	Value  *int
	Number struct {
		Value  Value
		Points []Point
	}
	Symbol struct {
		Symbol rune
		Point
	}
)

// returns all symbols in a line
func parseSymbols(line string, y int) []Symbol {
	symbols := make([]Symbol, 0)
	for x, ch := range line {
		if ch != '.' && !unicode.IsDigit(ch) {
			symbols = append(symbols, Symbol{Symbol: ch, Point: Point{X: x, Y: y}})
		}
	}
	return symbols
}

// returns all numbers in a line
func parseNums(line string, y int) []Number {
	var (
		numbers   = make([]Number, 0)
		currChars = ""
		currStart = 0
	)
	for i, ch := range line {
		if unicode.IsDigit(ch) {
			currChars += string(ch)
		} else {
			if currChars != "" {
				num := parsePoints(currChars, currStart, y)
				numbers = append(numbers, num)
				currChars = ""
			}
			currStart = i + 1
		}
	}

	// if we're at the end we have to make sure we append the last number if it exists
	if currChars != "" {
		num := parsePoints(currChars, currStart, y)
		numbers = append(numbers, num)
	}

	return numbers
}

// returns all possible points where the number.Value can be adjacent to a symbol
func parsePoints(num string, x, y int) Number {
	points := make([]Point, 0, len(num))
	for j := range num {
		points = append(points, Point{X: x + j, Y: y})
	}
	value := MustParse(num)
	return Number{Points: points, Value: &value}
}

func MustParse(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
