package main

import (
	"fmt"

	"honnef.co/go/js/dom"
)

func main() {
	// Here I code up the majority of today's blog entry in Go, meta!
	d := dom.GetWindow().Document()

	d.GetElementByID("blog-title").SetInnerHTML("Joseph's Golang Blog")
	d.GetElementByID("blog-subtitle").SetInnerHTML("Learning the Ways of the Gopher")

	d.GetElementByID("post-1-title").SetInnerHTML("Day 1: Setting up GopherJS")
	d.GetElementByID("post-1-attribution").SetInnerHTML("By Joseph Choi under <a class='post-category post-category-js'>Go</a>")
	d.GetElementByID("post-1-description-1").SetInnerHTML("Today was my first day with GopherJS. In order to drive myself to learn this new language, I am hoping to make short, Go-related posts on a regular basis. Wish me luck!")
	fmt.Println(d)
}
