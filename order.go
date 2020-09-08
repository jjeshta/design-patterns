package main

type order struct {
	cook *chef
	food string
}

func newOrder(cook *chef, food string) order {
	return order{cook, food}
}

func (o *order) execute() {
	if o.food == "pasta" {
		o.cook.cookPasta()
	} else {
		o.cook.bakeCake()
	}
}
