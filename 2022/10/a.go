package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveA(input string) string {
	var (
		sc  = bufio.NewScanner(strings.NewReader(input))
		cpu = NewCPU()
	)
	for sc.Scan() {
		cpu.AddInstruction(sc.Text())
	}
	cpu.Run(false)
	return fmt.Sprint(cpu.SignalStrength())
}
