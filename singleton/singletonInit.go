package singleton

import "fmt"

type a struct {
	str string
}

var singletonA *a

func init() {
	//initialize static instance on load
	fmt.Println("Initializing singleton A")
	singletonA = &a{str: "abc"}
}

//GetInstanceA - get singleton instance pre-initialized
func GetInstanceA() *a {
	return singletonA
}
