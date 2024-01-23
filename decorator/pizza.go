package main

type IPizza interface {
	getPrice() int
}

type VeggiePizza struct{}

func (pizza VeggiePizza) getPrice() int {
	return 60
}
