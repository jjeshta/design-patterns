package main

type burgerBuilder interface {
	setPatty(string)
	setTomato(int)
	getFood() burger
	String() string
}

func getBuilder(builderType string) burgerBuilder {
	if builderType == "veg" {
		return newVegBuilder()
	}
	if builderType == "non-veg" {
		return newMeatBuilder()
	}
	return nil
}
