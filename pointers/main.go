package main

import "fmt"

func main() {
	myNum := 7

	thala := &myNum
	fmt.Println("your num address is ", thala)
	fmt.Println("your num  is ", *thala)
	*thala = *thala * 2

	fmt.Println("your new num is ", *thala)

}
