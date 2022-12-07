package main

import (
	"bufio"
	"strconv"
	"strings"
)

type NodeType uint8

const (
	NodeTypeFolder NodeType = iota
	NodeTypeFile
)

type Node interface {
	Name() string
	Type() NodeType
	Size() int
	Parent() *FolderNode
}

func ScanLs(sc *bufio.Scanner) []string {
	lsOutput := make([]string, 0)
	for sc.Scan() {
		if strings.HasPrefix(sc.Text(), "$") {
			break
		}
		lsOutput = append(lsOutput, sc.Text())
	}
	return lsOutput
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

type FileNode struct {
	parent *FolderNode
	name   string
	sz     int
}

func (fn *FileNode) Name() string {
	return fn.name
}

func (fn *FileNode) Type() NodeType {
	return NodeTypeFile
}

func (fn *FileNode) Size() int {
	return fn.sz
}

func (fn *FileNode) Parent() *FolderNode {
	return fn.parent
}

type FolderNode struct {
	parent   *FolderNode
	name     string
	children []Node
}

func (fn *FolderNode) Name() string {
	return fn.name
}

func (fn *FolderNode) Type() NodeType {
	return NodeTypeFolder
}

func (fn *FolderNode) Size() int {
	totalSize := 0
	for _, v := range fn.children {
		totalSize += v.Size()
	}
	return totalSize
}

func (fn *FolderNode) Parent() *FolderNode {
	if fn.parent == nil {
		return fn
	}
	return fn.parent
}

func MustParseInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}
