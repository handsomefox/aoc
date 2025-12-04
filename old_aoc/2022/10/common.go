package main

import (
	"strconv"
	"strings"
)

type InstructionType byte

const (
	_ InstructionType = iota
	InstructionNoop
	InstructionAddx
)

type Instruction struct {
	f   func(cpu *CPU)
	typ InstructionType
}

type ScheduledInstruction struct {
	*Instruction
	waitCycles int
}

type CPU struct {
	picture string

	scheduled    *ScheduledInstruction
	instructions []Instruction

	X          int
	cycleCount int
	signalSum  int
}

func NewCPU() *CPU {
	return &CPU{
		X: 1,
	}
}

func (cpu *CPU) Picture() string {
	return cpu.picture
}

func (cpu *CPU) SignalStrength() int {
	return cpu.signalSum
}

func (cpu *CPU) AddInstruction(istr string) {
	if strings.HasPrefix(istr, "noop") {
		cpu.instructions = append(cpu.instructions, Instruction{
			f:   func(cpu *CPU) {},
			typ: InstructionNoop,
		})
		return
	}
	if strings.HasPrefix(istr, "addx") {
		cpu.instructions = append(cpu.instructions, Instruction{
			f:   func(cpu *CPU) { cpu.X += MustParseInt(strings.Split(istr, " ")[1]) },
			typ: InstructionAddx,
		})
		return
	}
}

func (cpu *CPU) Run(shouldDraw bool) {
	for {
		switch cpu.cycleCount {
		case 20, 60, 100, 140, 180, 220:
			cpu.signalSum += cpu.cycleCount * cpu.X
		}

		if cpu.hasScheduled() {
			ran := cpu.runScheduled()
			if !ran {
				if shouldDraw {
					cpu.draw()
				}
				cpu.cycleCount++
				continue
			}
		}

		if !cpu.shouldContinue() {
			return
		}

		i := cpu.next()
		switch i.typ {
		case InstructionAddx:
			cpu.schedule(i, 1)
		case InstructionNoop:
			cpu.schedule(i, 0)
		}

		if shouldDraw {
			cpu.draw()
		}
		cpu.cycleCount++
	}
}

func (cpu *CPU) draw() {
	i := cpu.cycleCount % 40
	x := cpu.X
	if i == x-1 || i == x || i == x+1 {
		cpu.picture += "#"
	} else {
		cpu.picture += "."
	}
	if i == 39 {
		cpu.picture += "\n"
	}
}

func (cpu *CPU) next() *Instruction {
	i := &cpu.instructions[0]
	cpu.instructions = cpu.instructions[1:]
	return i
}

func (cpu *CPU) hasScheduled() bool {
	return cpu.scheduled != nil
}

func (cpu *CPU) shouldContinue() bool {
	return len(cpu.instructions) != 0 || cpu.scheduled != nil
}

func (cpu *CPU) schedule(i *Instruction, count int) {
	cpu.scheduled = &ScheduledInstruction{
		Instruction: i,
		waitCycles:  count,
	}
}

func (cpu *CPU) runScheduled() bool {
	if cpu.scheduled.waitCycles == 0 {
		cpu.scheduled.f(cpu)
		cpu.scheduled = nil
		return true
	}

	cpu.scheduled.waitCycles--
	return false
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
