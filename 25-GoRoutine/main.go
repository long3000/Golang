package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
		"http://youtube.com",
		"http://taobao.com",
	}

	for _, link := range links {
		requestChecker(link)
	}
}

func requestChecker(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!")
		return // Don't do anything else
	}

	fmt.Println(link, " is up!")
}
