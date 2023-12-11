package main

import (
	"bufio"
	"slices"
	"strconv"
	"strings"
)

var (
	GoUp   = []byte{'S', '|', 'J', 'L'}
	RecvUp = []byte{'|', 'F', '7'}

	GoDown   = []byte{'S', '|', '7', 'F'}
	RecvDown = []byte{'|', 'J', 'L'}

	GoLeft   = []byte{'S', '-', 'J', '7'}
	RecvLeft = []byte{'-', 'L', 'F'}

	GoRight   = []byte{'S', '-', 'L', 'F'}
	RecvRight = []byte{'-', 'J', '7'}
)

func CanGoAndReceive(ch byte, canGo []byte, recvCh byte, canRecv []byte) bool {
	return slices.Contains(canGo, ch) && slices.Contains(canRecv, recvCh)
}

func findStart(grid [][]byte) (r, c int) {
	for r, row := range grid {
		for c, ch := range row {
			if ch == 'S' {
				return r, c
			}
		}
	}
	return 0, 0
}

type Coords struct {
	Row, Col int
}

type Deque[T any] struct {
	data []T
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{data: make([]T, 0)}
}

func (d *Deque[T]) Empty() bool {
	return len(d.data) == 0
}

func (d *Deque[T]) PushLeft(item T) {
	d.data = append([]T{item}, d.data...)
}

func (d *Deque[T]) PushRight(item T) {
	d.data = append(d.data, item)
}

func (d *Deque[T]) PopLeft() T {
	item := d.data[0]
	d.data = d.data[1:]
	return item
}

func (d *Deque[T]) PopRight() T {
	item := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return item
}

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

func Bytes(input string) [][]byte {
	sc := bufio.NewScanner(strings.NewReader(input))
	lines := make([][]byte, 0)
	for sc.Scan() {
		lines = append(lines, []byte(sc.Text()))
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
