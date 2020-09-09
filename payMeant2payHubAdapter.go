package main

//Adapter |implements the interface iPayHub ,
// payMeant2payHubAdapter adapts to payMeant
type payMeant2payHubAdapter struct {
	payObj payMeant
}

//The adapter pattern essentially wraps the old class's methods with the new class's method's names, and does so by:
//Implementing the old class's interface for the names of the adapter's methods.
//Holding a reference to the new class, and using the new class's methods as the adapter's methods bodies.
func (a *payMeant2payHubAdapter) addItem(itemName string) {
	a.payObj.addOneItem(itemName)
}

func (a *payMeant2payHubAdapter) addPrice(itemPrice string) {
	a.payObj.addPriceToTotal(itemPrice)
}

func newAdapter(payAdapter payMeant) *payMeant2payHubAdapter {
	return &payMeant2payHubAdapter{payAdapter}
}
