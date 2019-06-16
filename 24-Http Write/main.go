package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var url = "http://google.com"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, response.Body)
}
