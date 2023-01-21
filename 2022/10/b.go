package main

import (
	"bufio"
	"strings"
)

func SolveB(input string) string {
	var (
		sc  = bufio.NewScanner(strings.NewReader(input))
		cpu = NewCPU()
	)
	for sc.Scan() {
		cpu.AddInstruction(sc.Text())
	}
	cpu.Run(true)
	return "\n" + cpu.Picture()
}
