package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Lanternfish struct {
	currentTimer int
}

func passDay(list []Lanternfish) []Lanternfish {
	length := len(list)
	for i := 0; i < length; i++ {
		fish := &list[i]
		if fish.currentTimer == 0 {
			fish.currentTimer = 6
			list = append(list, fish.Spawn())
			continue
		}
		fish.currentTimer--
	}

	return list
}

func (f *Lanternfish) Spawn() Lanternfish {
	return Lanternfish{
		currentTimer: f.currentTimer + 2,
	}
}

func SolveA(input string) string {
	sn := bufio.NewScanner(strings.NewReader(input))

	list := make([]Lanternfish, 0)

	for sn.Scan() {
		str := sn.Text()

		values := strings.Split(str, ",")

		for _, value := range values {
			v, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			fish := Lanternfish{
				v,
			}

			list = append(list, fish)
		}
	}

	for i := 0; i < 80; i++ {
		list = passDay(list)
	}

	return fmt.Sprint(len(list))
}
