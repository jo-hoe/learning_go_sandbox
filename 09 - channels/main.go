package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://google.com",
	"https://facebook.com",
	"https://stackoverflow.com",
	"https://amazon.com",
	"https://golang.org",
}

func main() {
	c := make(chan string)

	for _, url := range urls {
		go avaiablityCheck(url, c)
	}

	for url := range c {
		// this is called a "function literal"
		// which is simliar to an anonymous function aka lambda
		go func(url string) {
			time.Sleep(time.Second * 2)
			go avaiablityCheck(url, c)
		}(url)
	}
}

// Note that this avaiability check considers a url to be up even if an
// http error code is returned. Only if a connection error occurs
// (e.g. DNS name cannot be resolved) the website is considered down.
func avaiablityCheck(url string, c chan string) {
	_, err := http.Get(url)

	if err == nil {
		fmt.Println(url, "- is up")
	} else {
		fmt.Println(url, "- cannot be reached")
	}
	c <- url
}
