package main

//This class is an Adaptee
import (
	"fmt"
)

type payMeant struct{}

func (p *payMeant) addOneItem(name string) {
	fmt.Printf("\n1 item added: %s\n", name)
}

func (p *payMeant) addPriceToTotal(price string) {
	fmt.Printf("1 item added to total with the price of: %s\n", price)
}

// Unique method
func (p *payMeant) addItemAndPrice(name, price string) {
	p.addOneItem(name)
	p.addPriceToTotal(price)
}
