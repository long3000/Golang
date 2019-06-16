package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var url = "http://google.com"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	byteSlice := make([]byte, 99999) // Make a byte slice with 99999 elements
	response.Body.Read(byteSlice)
	fmt.Println(string(byteSlice))
}
