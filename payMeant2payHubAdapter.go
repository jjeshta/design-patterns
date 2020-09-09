package main

type payMeant2payHubAdapter struct {
	payObj payMeant
}

func (a *payMeant2payHubAdapter) addItem(itemName string) {
	a.payObj.addOneItem(itemName)
}

func (a *payMeant2payHubAdapter) addPrice(itemPrice string) {
	a.payObj.addPriceToTotal(itemPrice)
}

func newAdapter(payAdapter payMeant) *payMeant2payHubAdapter {
	return &payMeant2payHubAdapter{payAdapter}
}
