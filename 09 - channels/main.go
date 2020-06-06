package main

import (
	"fmt"
	"net/http"
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

	fmt.Println(<-c)
}

// Note that this avaiability check considers a url to be up even if an
// http error code is returned. Only if a connection error occurs
// (e.g. DNS name cannot be resolved) the website is considered down.
func avaiablityCheck(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- url + "- cannot be reached"
		return
	}

	c <- url + "- is up"
}
