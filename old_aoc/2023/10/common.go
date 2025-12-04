package main

import (
	"bufio"
	"slices"
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

func Bytes(input string) [][]byte {
	sc := bufio.NewScanner(strings.NewReader(input))
	lines := make([][]byte, 0)
	for sc.Scan() {
		lines = append(lines, []byte(sc.Text()))
	}
	return lines
}
