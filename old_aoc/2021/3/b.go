package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolveB(in string) string {
	input := make([]string, 0, 1000)

	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Get oxygen generator rating
	ogr := findOxygenRating(input)
	fmt.Printf("Found oxygen generator rating: %v\n", ogr)

	// Get C02 scrubber rating
	c02 := findC02Rating(input)
	fmt.Printf("Found C02 scrubber rating: %v\n", c02)

	// Verify life support rating
	lsr := c02 * ogr
	return fmt.Sprintf("Found life support rating: %v", lsr)
}

func findOxygenRating(input []string) int64 {
	ogr := make([]string, 0)
	ogr = append(ogr, input...)
	for i := 0; i < len(input[0]); i++ {
		res := make([]string, 0)
		zero, one := countNumbers(i, ogr)

		if one >= zero {
			for _, v := range ogr {
				if string(v[i]) == "1" {
					res = append(res, v)
				}
			}
		} else {
			for _, v := range ogr {
				if string(v[i]) == "0" {
					res = append(res, v)
				}
			}
		}
		ogr = res
	}
	oxygenRating, _ := strconv.ParseInt(ogr[0], 2, 64)
	return oxygenRating
}

func findC02Rating(input []string) int64 {
	co2 := make([]string, 0)
	co2 = append(co2, input...)

	for i := 0; i < len(input[0]); i++ {
		res := make([]string, 0)
		zero, one := countNumbers(i, co2)

		if one >= zero {
			for _, v := range co2 {
				if string(v[i]) == "0" {
					res = append(res, v)
				}
			}
		} else {
			for _, v := range co2 {
				if string(v[i]) == "1" {
					res = append(res, v)
				}
			}
		}
		co2 = res
		if len(co2) == 1 {
			break
		}
	}

	c02Rating, _ := strconv.ParseInt(co2[0], 2, 64)
	return c02Rating
}

func countNumbers(idx int, input []string) (zero, one int) {
	for _, v := range input {
		if string(v[idx]) == "0" {
			zero++
		} else {
			one++
		}
	}
	return
}
