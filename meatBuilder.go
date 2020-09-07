package main

import "fmt"

type meatBuilder struct {
	patty  string
	tomato int
}

func newMeatBuilder() *meatBuilder {
	return &meatBuilder{}
}

func (b meatBuilder) String() string {
	return fmt.Sprintf("Your meat burger is Ready with the following options.\n Patty: %v, Tomato: %d", b.patty, b.tomato)
}

func (b *meatBuilder) setPatty(p string) {
	b.patty = p
}

func (b *meatBuilder) setTomato(t int) {
	b.tomato = t
}

func (b *meatBuilder) getBurger() burger {
	return burger{
		patty:  b.patty,
		tomato: b.tomato,
	}
}
