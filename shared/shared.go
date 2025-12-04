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
