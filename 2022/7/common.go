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
