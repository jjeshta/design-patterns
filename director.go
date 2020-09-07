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

func (d *director) buildBurger() burger {
	d.builder.setPatty()
	d.builder.setTomato()
	return d.builder.getBurger()
}
