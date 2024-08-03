package main

import "fmt"

func main() {
	// map work with key value pairs it just like maps in js or almost like objects in js
	var myMap = make(map[string]string)
	myMap["py"] = "python"
	myMap["js"] = "javascript"
	myMap["rb"] = "ruby"
	myMap["rs"] = "Rust"

	fmt.Println("myMaps", myMap)
	fmt.Println("here is Js", myMap["js"])

	// maps with loops

	for key, value := range myMap {
		fmt.Println("here is key value", key, value)

	}
}
