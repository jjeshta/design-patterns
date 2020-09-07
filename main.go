package main

import "fmt"

func main() {
	vegBuilder := getBuilder("veg")
	meatBuilder := getBuilder("non-veg")

	director := newDirector(vegBuilder)
	veggie := director.buildBurger("Mixed Veg", 4)

	fmt.Printf("Veg burger patty: %s\n", veggie.patty)
	fmt.Printf("Veg burger tomato: %d\n", veggie.tomato)

	director.setBuilder(meatBuilder)
	meat := director.buildBurger("Chicken", 2)

	fmt.Printf("\nNon Veg burger patty: %s\n", meat.patty)
	fmt.Printf("Non Veg burger tomato: %d\n", meat.tomato)
}
