package main

type ipayHub interface {
	addItem(itemName string)
	addPrice(itemPrice string)
}

// type ipayMeant interface {
// 	addItemAndPrice(name, price string)
// }
