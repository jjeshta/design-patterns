package main

import (
	"fmt"
)

type codeTwit struct {
	status string
	url    string
}

func (c *codeTwit) tweet() {
	fmt.Printf("Tweeted: %v from %v \n", c.status, c.url)
}
