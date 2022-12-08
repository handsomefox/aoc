package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func SolveB(input string) string {
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
	scores := make([]int, 0)
	for y, patch := range patches {
		for x := range patch {
			scores = append(scores, IsTreeVisible2(x, y, patches))
		}
	}
	sort.IntSlice(scores).Sort()
	return fmt.Sprint(scores[len(scores)-1]) // result is the best possible scenic score.
}

func IsTreeVisible2(treeX, treeY int, patches [][]int) int {
	if treeX == 0 || treeX == len(patches[0])-1 {
		return 0
	}
	if treeY == 0 || treeY == len(patches)-1 {
		return 0
	}
	return walk2(treeX, treeY, patches)
}

func walk2(treeX, treeY int, patches [][]int) int {
	// 4 directions to walk
	// we have to count the amount of trees
	// you can see in each direction.
	tree := patches[treeY][treeX]
	visibleTrees := make([]int, 0, 4) // amount of visible trees in each direction
	for i := 0; i < 4; i++ {
		count := 0
		treeX, treeY := treeX, treeY // copy starting coordinates for the iterations
		for !IsOnEdge(treeX, treeY, patches) {
			treeX, treeY = MustMoveToDirection(treeX, treeY, i)
			current := patches[treeY][treeX]
			if tree <= current {
				count++
				break
			}
			count++
		}
		visibleTrees = append(visibleTrees, count)
	}
	scenicScore := 1
	for _, c := range visibleTrees {
		scenicScore *= c
	}
	return scenicScore
}
