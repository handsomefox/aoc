package main

import (
	"bufio"
	"fmt"
	"strings"
)

func SolveA(input string) string {
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
	return fmt.Sprint(folderSize(rootNode))
}

func folderSize(root *FolderNode) int {
	folders := findEndFolders(root)
	size := 0
	for _, f := range folders {
		if f.Size() <= 100000 {
			size += f.Size()
		}
		for node := f.parent; node != root; node = node.parent {
			if node.Size() <= 100000 {
				size += node.Size()
			}
		}
	}
	return size
}

// findEndFolders returns a slice of folders that have no other folders inside of them.
func findEndFolders(root *FolderNode) []*FolderNode {
	var (
		folderCount = 0
		folders     = make([]*FolderNode, 0)
	)
	for _, child := range root.children {
		if child.Type() != NodeTypeFolder {
			continue
		}
		f, _ := child.(*FolderNode)
		folders = append(folders, findEndFolders(f)...)
		folderCount++
	}
	if folderCount == 0 {
		folders = append(folders, root)
	}
	return folders
}
