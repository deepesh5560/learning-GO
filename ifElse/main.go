package main

import "fmt"

func main() {
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
}
