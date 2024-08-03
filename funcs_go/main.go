package main

import "fmt"

func main() {
	fmt.Println("working on functions")

	// functions in GO works same as it works in js but u can not define a function inside a function

	//normal function
	greet()

	//function with parameters
	fmt.Println(add(2, 2))

	//function with multiple parameters
	fmt.Println(superAdder(23, 5, 7, 40))

}

func greet() {
	fmt.Println("hello world")
}

//in function with parameters the type of params should be defined and if function is returning something that type should also be defined

// here int with a and b are types of params and 3rd int outside round brackets is for return value type just like TS
func add(a int, b int) int {
	return a + b
}

// this is  how can we handle multiple params values same concept as JS
func superAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total = total + v
	}
	return total
}
