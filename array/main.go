package main

import "fmt"

func main() {
	//not used much but we can use these for fixed length data
	var numList [4]int

	numList[0] = 10
	numList[1] = 20
	numList[3] = 40

	fmt.Println("Array elements:", numList)
}
