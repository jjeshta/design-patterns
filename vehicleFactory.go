package main

import (
	"fmt"
)

func getVehicle(gunType string) (iVehicle, error) {
	if gunType == "car" {
		return newCar(), nil
	}
	if gunType == "ship" {
		return newShip(), nil
	}
	return nil, fmt.Errorf("Wrong vehicle type passed")
}
