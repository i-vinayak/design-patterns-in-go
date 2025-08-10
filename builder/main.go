package main

import "fmt"

func main() {
	builder := NewBuilder("wooden")

	builder.SetDoorType()
	builder.SetNumFloor(3)
	builder.SetWindowType()
	house := builder.GetHouse()

	fmt.Println(house.doorType, "\t", house.windowType, "\t", house.floors)
}
