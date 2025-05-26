package main

import "fmt"

func main() {
	repository := NewRepository()

	emptyHandler := NewEmptyHandler()
	lengthHandler := NewLengthHandler()
	repositoryHandler := NewRepositoryHandler(repository)

	// chaining the handlers
	emptyHandler.SetNext(lengthHandler).SetNext(repositoryHandler)

	// setting data
	err := emptyHandler.Handle("Hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Succesffuly set data")
	}

	fmt.Println(repository.GetData())
}
