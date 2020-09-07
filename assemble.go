package main

//assemble burger
type assemble struct{}

// assembleBurger Build a concrete  via burgerBuilder
func (a *assemble) assembleBurger(builder burgerBuilder) *burger {
	builder.setTomato()
	builder.setPatty()
	b := builder.getBurger()

	return &b
}
