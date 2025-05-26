package main

import "fmt"

func main() {
	repository := NewRepository()

	// injecting the repo layer into service
	s := NewService(repository)

	s.ValidateAndSet("random Data")
	fmt.Println(s.GetData())
}
