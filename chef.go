package main

import (
	"fmt"
)

type chef struct{}

func (c *chef) cookPasta() {
	fmt.Println("Chef is cooking Chicken Alfredo")
}

func (c *chef) bakeCake() {
	fmt.Println("Chef is baking Blueberry cheesecake")
}
