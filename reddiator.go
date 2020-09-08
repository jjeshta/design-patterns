package main

import (
	"fmt"
)

type reddiator struct {
	title string
	url   string
}

func (r *reddiator) reddit() {
	fmt.Printf("Reddit! url: %v from %v \n", r.title, r.url)
}
