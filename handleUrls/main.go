package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://go.dev:3000/packages?print=fmt&api=http"

func main() {
	fmt.Println("")

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Println("Protocol:", result.Scheme)   // http || https
	fmt.Println("Protocol:", result.Host)     // host name
	fmt.Println("Protocol:", result.Path)     //route
	fmt.Println("Protocol:", result.Port())   // port
	fmt.Println("Protocol:", result.RawQuery) //query parameters

	qPrams := result.Query() // its giving key value pair in map for all of the query parameters

	fmt.Printf("%T\nQuery Parameters: %v", qPrams, qPrams)

	// how to create a url

	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "LCO.dev",
		Path:     "packages",
		RawQuery: "user=Deep&&package=fmt",
	}

	newUrl := partsOfUrl.String()

	fmt.Println(newUrl)

}
