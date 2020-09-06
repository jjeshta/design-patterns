package main

import (
	"fmt"
	"patterns/singleton"
)

func main() {
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
