package common

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
	Float interface {
		~float32 | ~float64
	}
	Number interface {
		Integer | Float
	}
	Ordered interface {
		Integer | Float | ~string
	}

	MapFn[T any, R any]  func(value T) R
	AccumulatorFn[T any] func(total T, current T) T
	FilterFn[T any]      func(value T) bool

	Zipped[T any] struct {
		First, Second T
	}
)

func Int[T Signed](s string) T {
	return T(Must2(strconv.ParseInt(s, 10, 64)))
}

func Lines(input string) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		sc := bufio.NewScanner(strings.NewReader(input))
		for sc.Scan() {
			text := sc.Text()
			if text == "" {
				continue
			}
			c <- text
		}
	}()

	return c
}

func All[T any](values []T, condition func(T) bool) bool {
	for i := range values {
		if !condition(values[i]) {
			return false
		}
	}
	return true
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

func ForEach[T any](values []T, f func(T)) {
	for i := range values {
		f(values[i])
	}
}

func Sum[T Number](s []T) (sum T) {
	for _, i := range s {
		sum += i
	}
	return sum
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Map[T any, R any](values []T, mapper MapFn[T, R]) []R {
	mapped := make([]R, 0)
	for i := range values {
		mapped = append(mapped, mapper(values[i]))
	}
	return mapped
}

func Reduce[T any](values []T, accumulator AccumulatorFn[T], initialValue T) T {
	res := initialValue
	for i := range values {
		res = accumulator(res, values[i])
	}
	return res
}

func Filter[T any](values []T, filter FilterFn[T]) []T {
	res := make([]T, 0)
	for i := range values {
		if filter(values[i]) {
			res = append(res, values[i])
		}
	}
	return res
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

func SliceCopy[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

func Instances[T comparable](s []T, v T) int {
	count := 0
	for _, vv := range s {
		if vv == v {
			count++
		}
	}
	return count
}

func Assert(condition bool, message ...string) {
	if !condition {
		panic(formatMessage(message...))
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must2[T any](v T, err error) T {
	Must(err)
	return v
}

func formatMessage(message ...string) string {
	if len(message) == 0 {
		return "assertion failed"
	}
	return "assertion failed, " + strings.Join(message, ", ")
}
