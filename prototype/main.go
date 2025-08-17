package main

import "fmt"

func main() {
	file1 := NewFile("file-1")
	file2 := NewFile("file-2")
	files := []Node{file1, file2}

	folder1 := NewFolder("folder-1", files...)

	folder1Clone := folder1.Clone()

	folder1.Print()
	fmt.Println()
	folder1Clone.Print()

	fmt.Println()
	folder2 := NewFolder("main_folder", folder1, folder1Clone)
	folder2.Print()

	fmt.Println()
	folder2.Clone().Print()
}
