package main

import "fmt"

func main() {
	db1, err := NewDatabase("relational")
	if err != nil {
		fmt.Printf("Failed to create relational db - error: %s", err.Error())
		return
	}
	fmt.Println(db1)

	db2, err := NewDatabase("nosql")
	if err != nil {
		fmt.Printf("Failed to create nosql db - error: %s", err.Error())
		return
	}
	fmt.Println(db2)

	db3, err := NewDatabase("")
	if err != nil {
		fmt.Printf("Failed to create relational db - error: %s", err.Error())
		return
	}
	fmt.Println(db3)

}
