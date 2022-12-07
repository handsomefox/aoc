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

func cd(input string, current *FolderNode) *FolderNode {
	input = input[5:]
	switch input {
	// "/" is a special case, because it is the first command encountered and the folder
	// doesn't exist at this point.
	case "/":
		node := &FolderNode{
			parent:   current,
			name:     input,
			children: make([]Node, 0),
		}
		current.children = append(current.children, node)
		return node
	case "..":
		return current.Parent()
	default:
		for _, v := range current.children { // we find the specified folder inside the current one.
			if v.Name() == input && v.Type() == NodeTypeFolder {
				return v.(*FolderNode)
			}
		}
	}
	panic("unreachable") // unless the input has a mistake, this is unreachable, assume it doesn't
}

func ls(input []string, current *FolderNode) {
	for _, str := range input {
		if strings.HasPrefix(str, "dir") {
			folder := parseFolder(str, current)
			current.children = append(current.children, folder)
			continue
		}
		file := parseFile(str, current)
		current.children = append(current.children, file)
	}
}

func parseFolder(input string, parent *FolderNode) *FolderNode {
	foldername := strings.Split(input, " ")[1]
	return &FolderNode{
		parent:   parent,
		name:     foldername,
		children: make([]Node, 0),
	}
}

func parseFile(input string, parent *FolderNode) *FileNode {
	split := strings.Split(input, " ")
	fileSize, fileName := MustParseInt(split[0]), split[1]
	return &FileNode{
		parent: parent,
		name:   fileName,
		sz:     fileSize,
	}
}
