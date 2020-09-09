package main

import "fmt"

func main() {

	//Concret
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
