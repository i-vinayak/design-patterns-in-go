package main

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor(int)
	GetHouse() House
}

func NewBuilder(t string) Builder {
	if t == "wooden" {
		return &WoodenBuilder{}
	}
	if t == "glass" {
		return &GlassBuilder{}
	}
	return nil
}
