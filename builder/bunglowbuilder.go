package main

type BunglowBuilder struct {
	carpetArea int
	floors     int
	doors      int
	windows    int
}

func newBunglowBuilder() *BunglowBuilder {
	return &BunglowBuilder{}
}

func (builder *BunglowBuilder) setCarpetArea(carpetArea int) {
	builder.carpetArea = carpetArea
}

func (builder *BunglowBuilder) setFloors(floors int) {
	builder.floors = floors
}

func (builder *BunglowBuilder) setDoors(doors int) {
	builder.doors = doors
}

func (builder *BunglowBuilder) setWindows(windows int) {
	builder.windows = windows
}

func (builder *BunglowBuilder) getHouse() House {
	return House{
		carpetArea: builder.carpetArea,
		floors:     builder.floors,
		doors:      builder.doors,
		windows:    builder.windows,
	}
}
