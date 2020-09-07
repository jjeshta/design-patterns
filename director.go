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

func (d *director) buildBurger(p string, t int) burger {
	d.builder.setPatty(p)
	d.builder.setTomato(t)
	return d.builder.getBurger()
}
