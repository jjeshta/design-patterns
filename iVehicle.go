package main

type iVehicle interface {
	setMake(make string)
	setModel(model string)
	setType(vehicleType string)
	getMake() string
	getModel() string
	getType() string
}
