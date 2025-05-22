package main

import "fmt"

type Database interface {
	GetData(query string) string
}

type RealDatabase struct{}

func (rd *RealDatabase) GetData(query string) string {
	fmt.Println("Called in real db")
	if query == "1" {
		return "first"
	}
	return "hm"
}

func NewDatabase() Database {
	return &RealDatabase{}
}
