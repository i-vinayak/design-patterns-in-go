package main

import "fmt"

func main() {
	db1, err := NewDatabase("relational")
	if err != nil {
		fmt.Printf("Failed to create relational db - error: %s", err.Error())
		return
	}
	fmt.Println(db1)
	db1.PutData("This is DB1")
	fmt.Println(db1.GetData())

	db2, err := NewDatabase("nosql")
	if err != nil {
		fmt.Printf("Failed to create nosql db - error: %s", err.Error())
		return
	}
	fmt.Println(db2)
	db2.PutData("This is DB2")
	fmt.Println(db2.GetData())

	db3, err := NewDatabase("")
	if err != nil {
		fmt.Printf("Failed to create relational db - error: %s", err.Error())
		return
	}
	fmt.Println(db3)
	db3.PutData("This is DB3")
	fmt.Println(db3.GetData())

}
