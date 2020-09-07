package main

type burgerBuilder interface {
	setPatty(string)
	setTomato(int)
	getBurger() burger
}

func getBuilder(builderType string) burgerBuilder {
	if builderType == "veg" {
		return &vegBuilder{}
	}
	if builderType == "non-veg" {
		return &meatBuilder{}
	}
	return nil
}
