package main

type WoodenBuilder struct {
	windowType string
	doorType   string
	floors     int
}

func (b *WoodenBuilder) SetWindowType() {
	b.windowType = "Wooden window"
}

func (b *WoodenBuilder) SetDoorType() {
	b.doorType = "Wooden door"
}

func (b *WoodenBuilder) SetNumFloor(floors int) {
	b.floors = floors
}

func (b *WoodenBuilder) GetHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floors:     b.floors,
	}
}
