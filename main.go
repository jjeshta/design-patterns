package main

import "fmt"

func main() {
	toyota, _ := getVehicle("car")
	ship, _ := getVehicle("ship")
	printDetails(toyota)
	printDetails(ship)
}

func printDetails(v iVehicle) {
	fmt.Printf("Vehicle: %s - %s", v.getMake(), v.getModel())
	fmt.Println()
	fmt.Printf("Type: %s", v.getType())
	fmt.Println()
}
