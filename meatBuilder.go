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
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
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
