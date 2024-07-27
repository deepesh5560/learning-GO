package main

import "fmt"

const LoginToken = "login"

func main() {
	var username string
	var isLoggedIn bool = false
	email := "test@example.com"
	username = "testUser"

	fmt.Println(username, isLoggedIn, email, LoginToken)
	fmt.Printf("var of is type %T\n %T\n %T\n", username, isLoggedIn, email)

}
