package main

import "fmt"

type vegBuilder struct {
	patty  string
	tomato int
}

func newVegBuilder() *vegBuilder {
	return &vegBuilder{}
}

func (b vegBuilder) String() string {
	return fmt.Sprintf("Your veg burger is Ready with the following options.\n Patty: %v, Tomato: %d", b.patty, b.tomato)
}

func (b *vegBuilder) setPatty(p string) {
	b.patty = p
}

func (b *vegBuilder) setTomato(t int) {
	b.tomato = t
}

func (b *vegBuilder) getFood() burger {
	return burger{
		patty:  b.patty,
		tomato: b.tomato,
	}
}
