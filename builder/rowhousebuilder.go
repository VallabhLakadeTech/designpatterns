package main

type RowHouseBuilder struct {
	carpetArea int
	floors     int
	doors      int
	windows    int
}

func newRowHouseBuilder() *RowHouseBuilder {
	return &RowHouseBuilder{}
}

func (house *RowHouseBuilder) setCarpetArea(carpetArea int) {
	house.carpetArea = carpetArea
}

func (house *RowHouseBuilder) setFloors(floors int) {
	house.floors = floors
}

func (house *RowHouseBuilder) setDoors(doors int) {
	house.doors = doors
}

func (house *RowHouseBuilder) setWindows(windows int) {
	house.windows = windows
}

func (house *RowHouseBuilder) getHouse() House {
	return House{
		carpetArea: house.carpetArea,
		floors:     house.floors,
		doors:      house.doors,
		windows:    house.windows,
	}
}
