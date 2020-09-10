package main

import (
	"fmt"
)

func getVehicle(vehicleType string) (iVehicle, error) {
	if vehicleType == "car" {
		return newCar(), nil
	}
	if vehicleType == "ship" {
		return newShip(), nil
	}
	return nil, fmt.Errorf("Wrong vehicle type passed")
}
