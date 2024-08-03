package main

import (
	"fmt"
	"sort"
)

func main() {

	//slices are almost like  arrays in js
	mySlices := []string{}

	fmt.Println("mySlices", mySlices)

	mySlices = append(mySlices, "mango", "apple", "orange", "peach", "banana") //this is used to just like push in js but we can multipe elements in one go
	fmt.Println("mySlices app", mySlices)
	mySlices = append(mySlices[:2], mySlices[2:5]...) //[:2],[2:5] these will work just like slice where left val starting index and right is ending index-1
	fmt.Println("mySlices app2", mySlices)

	// using make

	scores := make([]int, 4) //here we are telling that what should be the length and thats fixed but we can change that by using append method

	scores[0] = 70
	scores[1] = 60
	scores[2] = 80
	scores[3] = 90
	// scores[4] = 30 // this will give err

	fmt.Println("scores", scores)

	scores = append(scores, 60, 45, 75) // this will give re alloacate the memory

	fmt.Println("scores2", scores)

	// sorting

	sort.Ints(scores)
	fmt.Println("sorted scores", scores)
	fmt.Println(" scores is Sorted?", sort.IntsAreSorted(scores))

}
