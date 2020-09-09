package main

import (
	"fmt"
	"patterns/singleton"
)

func main() {
	x := singleton.GetInstanceA()
	y := singleton.GetInstanceA()
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println()
	
	fmt.Println("Using sync.Once()")
	//First creation
	stateManager := singleton.GetManager()
	if stateManager.GetState() == "off" {
		stateManager.SetState("on")
	}
	fmt.Println(stateManager)

	//second creation
	stateManager1 := singleton.GetManager()
	if stateManager1.GetState() == "on" {
		stateManager1.SetState("off")
	}
	fmt.Println(stateManager1)

}
