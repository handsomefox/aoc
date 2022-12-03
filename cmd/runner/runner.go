package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
)

var (
	year string
	day  string
)

func init() {
	var (
		y = flag.String("y", "2022", "Year to run")
		d = flag.String("d", "1", "Day to run")
	)
	flag.Parse()
	year, day = *y, *d
}

func main() {
	cmd := exec.Command("go", "run", ".")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmd.Dir = path.Join(wd, year, day)

	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
