package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Starting")

	rand.Seed(time.Now().UnixNano())

	diceNum := rand.Intn(6) + 1

	fmt.Printf("You rolled a %d\n", diceNum)

	switch diceNum { //just liek js
	case 1:
		fmt.Printf("You can move %v\n", diceNum)

	case 2:
		fmt.Printf("You can move %v\n", diceNum)
		//  fallthrough  <= this is used id u want to run case next to it also
	case 3:
		fmt.Printf("You can move %v\n", diceNum)
	case 4:
		fmt.Printf("You can move %v\n", diceNum)
	case 5:
		fmt.Printf("You can move %v\n", diceNum)
	case 6:
		fmt.Printf("You can move %v\n", diceNum)
	default:
		fmt.Printf("%v this value is not possible try again\n", diceNum)

	}
}
