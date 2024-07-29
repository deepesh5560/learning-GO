package main

import "fmt"

func main() {
	// defer is a keyword when we use it it pushes that piece of code in the end of function
	// if there are multiple defer in the function it will be executed in reverse order (last in first out)
	defer fmt.Println("hello")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("World")
	mydefer()
} //result world 0 1 2 3 4 two one hello

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
