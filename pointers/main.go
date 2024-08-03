package main

import "fmt"

func main() {

	// nothing but refrence concept as we have in js and u have learned this concept in th basics of rust also (for me only)
	// here &myNum is a reference to specfic var it gives only address not copy but *myNum will give access to original value and we
	//change it diretly from  source
	myNum := 7

	thala := &myNum
	fmt.Println("your num address is ", thala)
	fmt.Println("your num  is ", *thala)
	*thala = *thala * 2

	fmt.Println("your new num is ", *thala)
	// or
	fmt.Println("your new num is ", myNum)

}
