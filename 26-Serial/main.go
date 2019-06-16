package main

import (
	"fmt"
	"net/http"
	"time"
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

	// Create new channel to handle STRING
	c := make(chan string)

	for _, link := range links {
		go checkRequest(link, c)
	}

	//
	for l := range c {
		// go checkRequest(l, c)
		// User function literal (aka. Lamda)
		go func() {
			time.Sleep(5 * time.Second)
			checkRequest(l, c)
		}()
	}
}

func checkRequest(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be dead ...")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
