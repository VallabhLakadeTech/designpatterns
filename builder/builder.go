package main

type IBuilder interface {
	setFloors(int)
	setCarpetArea(int)
	setDoors(int)
	setWindows(int)
	getHouse() House
}

func getBuilder(buildingType string) IBuilder {
	if buildingType == "rowhouse" {
		return newRowHouseBuilder()
	} else if buildingType == "bunglow" {
		return newBunglowBuilder()
	}
	return nil
}
