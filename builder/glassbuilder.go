package main

type GlassBuilder struct {
	windowType string
	doorType   string
	floors     int
}

func (b *GlassBuilder) SetWindowType() {
	b.windowType = "glass window"
}

func (b *GlassBuilder) SetDoorType() {
	b.doorType = "wood and glass door"
}

func (b *GlassBuilder) SetNumFloor(floors int) {
	b.floors = floors
}

func (b *GlassBuilder) GetHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floors:     b.floors,
	}
}
