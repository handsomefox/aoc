package main

import (
	"encoding/json"
	"strings"
)

type Pair struct {
	A, B []any
}

func ParseInput(input string) []Pair {
	var (
		unparsedPairs = strings.Split(input, "\n\n")
		pairs         = make([]Pair, 0)
	)
	for _, pair := range unparsedPairs {
		var (
			values = strings.Split(pair, "\n")
			a, b   = []any{}, []any{}
		)
		MustUnmarshal([]byte(values[0]), &a)
		MustUnmarshal([]byte(values[1]), &b)

		pairs = append(pairs, Pair{
			A: a,
			B: b,
		})
	}

	return pairs
}

// Compare returns:
//
//	-1, if left side is smaller,
//	 0, if both sides are equal,
//	 1, if left side is bigger.
func Compare(left, right any) int {
	lList, lIsList := left.([]any)
	rList, rIsList := right.([]any)

	if !lIsList && !rIsList {
		if left.(float64) < right.(float64) {
			return -1
		}
		if left.(float64) > right.(float64) {
			return 1
		}
		return 0
	}

	if lIsList && rIsList {
		min := Min(len(lList), len(rList))
		for i := 0; i < min; i++ {
			result := Compare(lList[i], rList[i])
			if result == -1 || result == 1 {
				return result
			}
		}
	}

	if lIsList && rIsList {
		if len(lList) < len(rList) {
			return -1
		} else if len(lList) > len(rList) {
			return 1
		}
	}

	if lIsList && !rIsList {
		list := []any{right}
		return Compare(lList, list)
	}

	if !lIsList && rIsList {
		list := []any{left}
		return Compare(list, rList)
	}

	return 0
}

func MustUnmarshal(data []byte, v any) {
	if err := json.Unmarshal(data, v); err != nil {
		panic(err)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
