package main

import "fmt"

func main() {

	deep := User{"deepesh", "deepesh@go.dev", 24, true}
	fmt.Println("User as deep", deep)
	fmt.Println("User name and mail", deep.Name, deep.Mail)

}

type User struct { // using first letter capiytal so we can make these public
	Name     string
	Mail     string
	Age      int
	IsActive bool
}
