package main

import (
	"bufio"
	"strconv"
	"strings"
)

type (
	Signed interface {
		~int | ~int64 | ~int32 | ~int16 | ~int8
	}
	Unsigned interface {
		~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8
	}
	Integer interface {
		Signed | Unsigned
	}

	MapFn[T any, R any] func(value T) R

	Zipped[T any] struct {
		First, Second T
	}
)

func All[T any](values []T, condition func(T) bool) bool {
	for i := range values {
		if !condition(values[i]) {
			return false
		}
	}
	return true
}

func Lines(input string) []string {
	sc := bufio.NewScanner(strings.NewReader(input))
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func MustParseInteger[T Integer](s string) T {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return T(i)
}

func MustParseIntegerSlice[T Integer](s string) []T {
	fields := strings.Fields(s)
	ints := make([]T, 0)
	for i := range fields {
		ints = append(ints, MustParseInteger[T](fields[i]))
	}
	return ints
}

func Map[T any, R any](values []T, mapper MapFn[T, R]) []R {
	mapped := make([]R, 0)
	for i := range values {
		mapped = append(mapped, mapper(values[i]))
	}
	return mapped
}

func Zip[T any](first, second []T) []Zipped[T] {
	length := min(len(first), len(second))
	zipped := make([]Zipped[T], length)

	for i := range zipped {
		zipped[i].First = first[i]
		zipped[i].Second = second[i]
	}

	return zipped
}
