package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveA(input string) string {
	var (
		sc      = bufio.NewScanner(strings.NewReader(input))
		patches = make([][]int, 0)
	)
	for sc.Scan() {
		var (
			text  = sc.Text()
			patch = make([]int, 0, len(text))
		)
		for _, r := range text {
			patch = append(patch, MustParseInt(string(r)))
		}
		patches = append(patches, patch)
	}
	count := 0
	for y, patch := range patches {
		for x := range patch {
			if IsTreeVisible(x, y, patches) {
				count++
			}
		}
	}
	return fmt.Sprint(count)
}

func IsTreeVisible(treeX, treeY int, patches [][]int) bool {
	if treeX == 0 || treeX == len(patches[0])-1 {
		return true
	}
	if treeY == 0 || treeY == len(patches)-1 {
		return true
	}
	return walk(treeX, treeY, patches)
}

func walk(treeX, treeY int, patches [][]int) bool {
	// 4 directions to walk
	// If the tree is visible in at least one of them,
	// return immediately.
	tree := patches[treeY][treeX]
	for i := 0; i < 4; i++ {
		failed := false
		treeX, treeY := treeX, treeY // copy starting coordinates for the iterations
		for !IsOnEdge(treeX, treeY, patches) {
			treeX, treeY = MustMoveToDirection(treeX, treeY, i)
			current := patches[treeY][treeX]
			if tree <= current {
				failed = true
				break
			}
		}
		if !failed {
			current := patches[treeY][treeX]
			if tree > current {
				return true
			}
		}
	}
	return false
}
