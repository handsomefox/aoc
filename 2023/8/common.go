package main

import "strings"

type Node struct {
	Left       *Node
	Right      *Node
	Identifier string // AAA, ZZZ
}

func ParseDirection(input string) string {
	split := strings.Split(input, "\n\n")
	return split[0]
}

func ParseNodes(input string) map[string]*Node {
	split := strings.Split(input, "\n\n")
	lines := strings.Split(split[1], "\n")

	// This is done in two passes.
	// 1. Fill the map with all the nodes from the left side of the " = ".
	nodes := make(map[string]*Node)
	for _, line := range lines {
		if line == "" {
			continue
		}
		ident := strings.Split(line, " =")[0]
		nodes[ident] = &Node{
			Identifier: ident,
		}
	}

	// 2. Add left and right fields to the filled nodes.
	for _, line := range lines {
		if line == "" {
			continue
		}
		split := strings.Split(line, "= ")

		children := strings.Fields(strings.Map(func(r rune) rune {
			switch r {
			case '(':
				return ' '
			case ')':
				return ' '
			case ',':
				return ' '
			default:
				return r
			}
		}, split[1]))

		node := strings.TrimSpace(split[0])
		nodes[node].Left = nodes[children[0]]
		nodes[node].Right = nodes[children[1]]
	}
	return nodes
}
