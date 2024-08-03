package main

import "fmt"

const LoginToken = "login" // global variable

// in go variables are block scope only if they are not defined globally

func main() {
	var username string
	var isLoggedIn bool = false
	email := "test@example.com" // := this is valrous operator in this we dont have define type and also no need of var key wor
	username = "testUser"

	{
		// test:="test"  now this variable can be used in this block only
	} // test variable will get dropped

	fmt.Println(username, isLoggedIn, email, LoginToken)
	fmt.Printf("var of is type %T\n %T\n %T\n", username, isLoggedIn, email)

}
