package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func SolveA(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	values := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		values = append(values, text)
	}

	if len(values) == 0 {
		log.Fatal(errors.New("empty input"))
	}

	bits := len(values[0]) // each line is this amount of bits (0,1)
	var (
		leastCommon string
		mostCommon  string
	)

	for i := 0; i < bits; i++ {
		most, _ := findMostCommonBit(values, i)
		least, _ := findLeastCommonBit(values, i)

		mostCommon += most
		leastCommon += least
	}

	mostDecimal, err := strconv.ParseInt(mostCommon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	leastDecimal, err := strconv.ParseInt(leastCommon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	lifeSupportRating := mostDecimal * leastDecimal
	return fmt.Sprintf("Life support rating: %d", lifeSupportRating)
}

func findMostCommonBit(values []string, pos int) (string, int) {
	var (
		zero int
		one  int
	)
	comp := byte('1')

	for _, v := range values {
		if v[pos] == comp {
			one++
		} else {
			zero++
		}
	}

	if one > zero || one == zero {
		return "1", one
	}
	return "0", zero
}

func findLeastCommonBit(values []string, pos int) (string, int) {
	var (
		zero int
		one  int
	)
	comp := byte('1')

	for _, v := range values {
		if v[pos] == comp {
			one++
		} else {
			zero++
		}
	}

	if zero > one || one == zero {
		return "0", zero
	}
	return "1", one
}
