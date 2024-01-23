package main

import "fmt"

func main() {

	veggiePizza := VeggiePizza{}
	fmt.Println("Normal veggie: ", veggiePizza.getPrice())

	tomatoTopping := TomatoTopping{pizza: veggiePizza}
	fmt.Println("Normal veggie with tomato topping: ", tomatoTopping.getPrice())

	cheeseTomatoTopping := CheeseTopping{pizza: tomatoTopping}
	fmt.Println("Normal veggie with tomato and cheese topping: ", cheeseTomatoTopping.getPrice())

}
