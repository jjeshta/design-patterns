package main

type meatBuilder struct {
	patty  string
	tomato int
}

func newMeatBuilder() *meatBuilder {
	return &meatBuilder{}
}

func (b *meatBuilder) setPatty() {
	b.patty = "Mutton"
}

func (b *meatBuilder) setTomato() {
	b.tomato = 1
}

func (b *meatBuilder) getBurger() burger {
	return burger{
		patty:  b.patty,
		tomato: b.tomato,
	}
}
