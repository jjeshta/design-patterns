package main

//Concrete Prototype   1 implements ipayHub interface
import (
	"fmt"
)

type payHub struct {
}

func (p *payHub) addItem(itemItem string) {
	fmt.Printf("\n1 item added: %s\n", itemItem)
}

func (p *payHub) addPrice(itemPrice string) {
	fmt.Printf("1 item added to total with the price of: %s\n", itemPrice)
}
