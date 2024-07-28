package main

import (
	"fmt"
)

func main() {

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// normal for loop
	for i := 0; i < len(days); i++ {
		fmt.Printf("%d. %s\n", i+1, days[i])
	}

	// for range loop with index its just like forEach loop in js but here ( i )will  give index instead of value
	for i := range days {
		fmt.Printf("%d. %s\n", i+1, days[i])
	}

	// this is just like forEach in js where i is index and day is value
	for i, day := range days {
		fmt.Printf("%d. %s\n", i+1, day)
	}

	// this is just like while loop in js

	index := 0

	// break is used to terminate the loop

	// where continue is used to skip 1 phase of it

	// goto label are kind of methods when we use them loop will terminate itself and execute that method
	for index < 10 {
		if index == 5 {

			fmt.Println("we found", index)
			index++

			continue
		}
		if index == 2 {
			goto deep

		}
		if index == 8 {
			fmt.Println("we found", index)
			break
		}
		fmt.Printf("%d\n", index)

		index++
	}

deep:
	fmt.Println("i m deep")

}
