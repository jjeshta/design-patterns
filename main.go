package main

import (
	"fmt"
)

func main() {
	withoutFacade()
	fmt.Println("---")
	withFacade()

}

func withoutFacade() {
	c := codeTwit{
		status: "Read my greatest meme ever!",
		url:    "https://meme-forever.com/post-meme",
	}

	g := googlize{
		url: "https://meme-forever.com/post-meme",
	}

	r := reddiator{
		title: "My greatest meme",
	}

	c.tweet()
	g.share()
	r.reddit()

}

func withFacade() {
	s := shareFacade{}
	s.newShareFacade("https://meme-forever.com/post-meme", "My greatest meme", "Read my greatest meme ever!")
	s.share()
}
