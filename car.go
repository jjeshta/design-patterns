package main

type car struct {
	vehicle
}

func newCar() iVehicle {
	return &car{
		vehicle: vehicle{
			make:        "Toyota",
			model:       "Corolla",
			vehicleType: "Motor Car",
		},
	}
}
