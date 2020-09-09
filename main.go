package main

import "fmt"

/**
Target: ipayHub interface defines two methods addPrice() and addItem()
Concrete Prototype: payHub (implements ipayHub) and payMeant2payHubAdapter
Adapter: payMeant2payHubAdapter
Adaptee: payMeant

Problem statement:
customer can buy item and pay with payHub as buy() takes in parameters iPayHub iterface.
therefore if customer buy items and pay with payMeant then errors will occur.
so we need to adapt payMeant as
payMeant offers the same functionality but through a different interface( addOneItem() and  addPriceToTotal())
**/
func main() {
	//Adaptee
	paymeant := payMeant{}
	payhub := &payHub{}

	//Adapter
	adapter := newAdapter(paymeant)

	//Client
	cust := customer{}
	cust.buy(payhub, "apple", "$5")

	fmt.Println("---")

	cust.buy(adapter, "apple", "$5")

}
