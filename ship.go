package main

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
