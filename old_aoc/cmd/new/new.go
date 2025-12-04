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

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func init() {
	var (
		y = flag.String("y", "2022", "Year to create")
		d = flag.String("d", "1", "Day to create")
	)
	flag.Parse()
	year, day = *y, *d
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	p := path.Join(wd, year, day)
	if ex, _ := exists(p); ex {
		fmt.Println("Path already exists, exiting")
		return
	}

	cmd := exec.Command("cp", "-r", "./template", p)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	b, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	if err := os.Remove(path.Join(p, "go.mod")); err != nil {
		panic(err)
	}

	cmd = exec.Command("go", "mod", "init", year+"_"+day)
	cmd.Dir = p
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	b, err = cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	cmd = exec.Command("go", "work", "use", "./"+year+"/"+day)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	b, err = cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
