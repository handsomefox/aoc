package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

func SolveB(input string) string {
	var (
		sc          = bufio.NewScanner(strings.NewReader(input))
		rootNode    = &FolderNode{parent: nil, name: "", children: make([]Node, 0)}
		currentNode = rootNode
	)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "$ ls") {
			ls(ScanLs(sc), currentNode)
		}
		if strings.HasPrefix(sc.Text(), "$ cd") {
			currentNode = cd(sc.Text(), currentNode)
		}
	}
	var (
		freeSpace   = 70000000 - rootNode.Size()
		spaceToFree = 30000000 - freeSpace
		candidates  = findPossible(rootNode, spaceToFree)
	)
	return fmt.Sprint(findMin(candidates))
}

func findMin(folders []*FolderNode) int {
	min := math.MaxInt
	for _, s := range folders {
		if s.Size() < min {
			min = s.Size()
		}
	}
	return min
}

func findPossible(root *FolderNode, spaceToFree int) []*FolderNode {
	suitable := make([]*FolderNode, 0)
	for _, child := range root.children {
		if child.Type() != NodeTypeFolder {
			continue
		}
		f, _ := (child.(*FolderNode))
		if child.Size() < spaceToFree {
			continue
		}
		suitable = append(suitable, f)                               // current child is a suitable candidate
		suitable = append(suitable, findPossible(f, spaceToFree)...) // still search for other ones
	}
	return suitable
}
