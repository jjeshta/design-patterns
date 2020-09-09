package main

// ship embedded vehicle struct and hence indirectly implements all methods from IVehicle interface
type ship struct {
	vehicle
}

func newShip() iVehicle {
	return &ship{
		vehicle: vehicle{
			make:        "Ocean 3",
			model:       "Cape",
			vehicleType: "Ship",
		},
	}
}
