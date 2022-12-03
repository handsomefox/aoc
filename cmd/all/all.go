package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	paths := make([]string, 0)

	filepath.Walk(wd, func(path string, _ fs.FileInfo, _ error) error {
		if strings.Contains(path, "main.go") && !strings.Contains(path, "template") {
			paths = append(paths, path[:len(path)-7])
		}
		return nil
	})

	for _, path := range paths {
		cmd := exec.Command("go", "run", ".")
		cmd.Dir = path
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		fmt.Println("Executing: ", path+"\\main.go")
		err := cmd.Run()
		if err != nil {
			fmt.Println("err: ", err)
		}
	}
}
