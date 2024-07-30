package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://api.restful-api.dev/objects"
const postUrl = "https://dummyjson.com/posts/add"

func main() {
	fmt.Println("")
	getApi()
	postApi()
}

func getApi() {

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

func postApi() {
	requestBody := strings.NewReader(`{
	"title":"I am in love with someone.",
	"userId":5
	}`) // creating JSON
	response, _ := http.Post(postUrl, "application/json", requestBody)
	result, _ := ioutil.ReadAll(response.Body)
	println("post", string(result))

	defer response.Body.Close()
}
