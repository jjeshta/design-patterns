package main

type customer struct {
}

func (c *customer) addItem(itemName string) {

}

func (c *customer) addPrice(itemPrice string) {

}

func (c *customer) buy(p ipayHub, itemName, itemPrice string) {
	p.addItem(itemName)
	p.addPrice(itemPrice)
}
