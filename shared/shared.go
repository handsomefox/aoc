package shared

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var UseSample bool = false

func init() {
	sample := flag.Bool("sample", false, "use sample.txt instead of input.txt")
	flag.Parse()
	UseSample = *sample
}

func MustReadInput(year, day int) string {
	filename := "input.txt"
	if UseSample {
		filename = "sample.txt"
	}
	rel := filepath.Join(strconv.Itoa(year), strconv.Itoa(day), filename)

	b, err := os.ReadFile(rel)
	if err != nil {
		panic(fmt.Errorf("read %s: %w", rel, err))
	}
	return strings.TrimRight(string(b), "\n")
}

func Lines(input string) <-chan string {
	sc := bufio.NewScanner(strings.NewReader(input))
	ch := make(chan string, 1)

	go func() {
		defer close(ch)
		for sc.Scan() {
			ch <- sc.Text()
		}
	}()

	return ch
}

func MustParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func String(v any) string {
	return fmt.Sprintf("%v", v)
}

func Map[T any, R any](in []T, f func(value T) R) []R {
	out := make([]R, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func Filter[T any](in []T, pred func(value T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T any, R any](in []T, acc R, f func(accumulator R, value T) R) R {
	for _, v := range in {
		acc = f(acc, v)
	}
	return acc
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CeilDiv(n, d int) int { return (n + d - 1) / d }
