package main

type TomatoTopping struct {
	pizza IPizza
}

func (topping TomatoTopping) getPrice() int {
	return topping.pizza.getPrice() + 20
}
