package main

import "fmt"

func main() { // just like js but it doesn't need round brackets to write conditional expressions
	num := 9

	if num > 10 {
		fmt.Println(num, "active user")

	} else if num < 10 {
		fmt.Println(num, "non active user")

	} else {
		fmt.Println(num, "Blocked user")
	}

	// declare var on the go

	if isActive := true; isActive {
		fmt.Println("User is active")
	} else {
		fmt.Println("User is not active")
	}

	test := ""

	if len(test) != 0 {
		fmt.Println("Condition is true")
	} else {
		fmt.Println("Condition is false")
	}

}
