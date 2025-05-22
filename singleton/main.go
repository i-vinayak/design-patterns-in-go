package main

import "fmt"

func main() {
	singleton := NewSingleton("First initialization")
	fmt.Println(singleton)
	fmt.Println(NewSingleton("Again"))
}
