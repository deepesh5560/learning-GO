package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://api.restful-api.dev/objects"

func main() {
	fmt.Println("")

	response, err := http.Get(url)

	if err != nil {
		panic(err)

	}

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)

	}

	result := string(dataBytes)

	fmt.Println("Result", result)

	defer response.Body.Close()
}
