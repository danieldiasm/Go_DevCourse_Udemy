package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, error := http.Get(link)
	if error != nil {
		fmt.Println(link, " might be down!")
		c <- "Might be down..."
		return
	}

	fmt.Println("Link ", link, " seems to be okay!")
	c <- "It is up!"
}
