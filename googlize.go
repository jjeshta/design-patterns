package main

import (
	"fmt"
)

type googlize struct {
	url string
}

func (g *googlize) share() {
	fmt.Printf("Shared on Google Plus: %s \n", g.url)
}
