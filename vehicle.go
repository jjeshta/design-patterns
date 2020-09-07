package main

type vehicle struct {
	make        string
	model       string
	vehicleType string
}

func (v *vehicle) setMake(make string) {
	v.make = make
}

func (v *vehicle) setModel(model string) {
	v.model = model
}

func (v *vehicle) setType(vehicleType string) {
	v.vehicleType = vehicleType
}

func (v *vehicle) getMake() string {
	return v.make
}

func (v *vehicle) getModel() string {
	return v.model
}

func (v *vehicle) getType() string {
	return v.vehicleType
}
