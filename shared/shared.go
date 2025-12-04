package shared

import (
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
