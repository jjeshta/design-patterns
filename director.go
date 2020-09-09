package main

type director struct {
	builder burgerBuilder
}

func newDirector(b burgerBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b burgerBuilder) {
	d.builder = b
}

// buildBurger method will get a builder interface and
// call the builder methods then return the object.
func (d *director) buildBurger(p string, t int) burger {
	d.builder.setPatty(p)
	d.builder.setTomato(t)
	return d.builder.getFood()
}
