package main

import "strconv"

func IsOnEdge(treeX, treeY int, patches [][]int) bool {
	if treeX == 0 || treeX == len(patches[0])-1 {
		return true
	}
	if treeY == 0 || treeY == len(patches)-1 {
		return true
	}
	return false
}

func MustMoveToDirection(treeX, treeY, direction int) (x, y int) {
	// .0.
	// 3P1
	// .2.
	switch direction {
	case 0:
		return treeX, treeY - 1
	case 1:
		return treeX + 1, treeY
	case 2:
		return treeX, treeY + 1
	case 3:
		return treeX - 1, treeY
	default:
		panic(direction)
	}
}

func MustParseInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
