package main

// What we need to implement is;
// An interface Command
// A class Order that implements Command interface
// A class Waiter (invoker)
// A class Chef (receiver)

func main() {
	//Client
	cook := chef{}
	placeOrder := newOrder(&cook, "pasta")
	attendant := newWaiter(&placeOrder)
	attendant.execute()

	placeOrder = newOrder(&cook, "cake")
	attendant = newWaiter(&placeOrder)
	attendant.execute()

}
