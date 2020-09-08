package main

type waiter struct {
	placeOrder *order
}

func newWaiter(placeOrder *order) waiter {
	return waiter{placeOrder}
}

func (w *waiter) execute() {
	w.placeOrder.execute()
}
