package main

import "fmt"

func main() {
	//getting instance of veg and nonveg burgerbuilders
	vegBuilder := getBuilder("veg")
	meatBuilder := getBuilder("non-veg")

	director := newDirector(vegBuilder)
	veggie := director.buildBurger("Mixed Veg", 4)
	fmt.Println(vegBuilder.String())

	director.setBuilder(meatBuilder)
	meat := director.buildBurger("Mutton", 2)
	fmt.Println(meatBuilder.String())

	fmt.Println("\n---------------------")
	fmt.Println(veggie)
	fmt.Println(meat)
}
