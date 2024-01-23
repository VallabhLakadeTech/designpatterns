package main

import "fmt"

func main() {
	fmt.Println("Welcome to builder pattern")

	house1 := getBuilder("rowhouse")
	house1.setCarpetArea(1350)
	house1.setDoors(6)
	house1.setFloors(2)
	house1.setWindows(5)
	fmt.Println(house1.getHouse())

	house2 := getBuilder("bunglow")
	house2.setCarpetArea(1750)
	house2.setDoors(14)
	house2.setFloors(3)
	house2.setWindows(10)
	fmt.Println(house2.getHouse())
}
