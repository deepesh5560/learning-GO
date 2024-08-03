package main

import "fmt"

func main() {
	//not used much but we can use these for fixed length data
	var numList [4]int // in go array length will  be fixed only these are not as dynamic as slices

	numList[0] = 10
	numList[1] = 20
	numList[3] = 40

	fmt.Println("Array elements:", numList)
}
