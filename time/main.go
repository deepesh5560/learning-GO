package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("")
	presentTime := time.Now()
	fmt.Println("Current date and time:", presentTime.Format("02-01-2006 15:04:05 Monday")) // always use this date only and we can change date format accordingly

	createdTime := time.Date(2000, time.July, 26, 5, 30, 0, 0, time.UTC) // month and timezone are always  fetched from time library
	fmt.Println("createdTime date and time:", createdTime.Format("02-01-2006 15:04:05 Monday"))

}
