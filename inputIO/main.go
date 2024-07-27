package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "wecoem here"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter rating for domino s:")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating,", input)

	var numRating, err = strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 ", numRating+1)
	}
}
