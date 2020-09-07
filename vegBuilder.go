package main

type vegBuilder struct {
	patty  string
	tomato int
}

func newVegBuilder() *vegBuilder {
	return &vegBuilder{}
}

func (b *vegBuilder) setPatty(p string) {
	b.patty = p
}

func (b *vegBuilder) setTomato(t int) {
	b.tomato = t
}

func (b *vegBuilder) getBurger() burger {
	return burger{
		patty:  b.patty,
		tomato: b.tomato,
	}
}
