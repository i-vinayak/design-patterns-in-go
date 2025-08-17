package main

import "fmt"

type Node interface {
	Clone() Node
	Print()
}

type file struct {
	name string
}

func (f *file) Clone() Node {
	return &file{name: f.name + "_clone"}
}

func (f *file) Print() {
	fmt.Println("File: ", f.name)
}

func NewFile(name string) Node {
	return &file{name: name}
}

type folder struct {
	name     string
	children []Node
}

func (f *folder) Clone() Node {
	clonedChildren := make([]Node, len(f.children))

	for i, child := range f.children {
		clonedChildren[i] = child.Clone()
	}

	return &folder{name: f.name + "_clone", children: clonedChildren}
}

func (f *folder) Print() {
	fmt.Println(f.name)
	for _, child := range f.children {
		fmt.Print("--- ")
		child.Print()
	}
}

func NewFolder(name string, children ...Node) Node {
	return &folder{name: name, children: children}
}
