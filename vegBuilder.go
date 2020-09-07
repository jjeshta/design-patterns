package main

type vegBuilder struct {
	patty  string
	tomato int
}

func newVegBuilder() *vegBuilder {
	return &vegBuilder{}
}

func (b *vegBuilder) setPatty() {
	b.patty = "Mixed Veg"
}

func (b *vegBuilder) setTomato() {
	b.tomato = 2
}

func (b *vegBuilder) getBurger() burger {
	return burger{
		patty:  b.patty,
		tomato: b.tomato,
	}
}
