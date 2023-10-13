package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	const url = "https://github.com"
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println(content)
}
