package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "working with the file system in GO...."

	file, errors := os.Create("./myFile.txt") //to create file

	checkError(errors)

	length, err := io.WriteString(file, content) // to write content to file

	checkError(err)

	fmt.Println("the lenth in", length)
	defer file.Close()
	readFile("./myFile.txt") // calling readFile
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename) //reading contents of file
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dataByte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	} //this function is used to check if error is not nil and then panic
}
