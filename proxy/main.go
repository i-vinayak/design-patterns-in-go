package main

import "fmt"

func main() {
	dbService := NewDatabaseService()
	fmt.Println(dbService.GetData("1"))
	fmt.Println(dbService.GetData("1"))
	fmt.Println(dbService.GetData("random query"))
}
