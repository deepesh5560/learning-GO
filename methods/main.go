package main

import "fmt"

func main() {
	deepesh := User{"Deepesh", "Deepesh@dev.go", 24, true}
	fmt.Println(deepesh.SetStatus(false))
	fmt.Println(deepesh.GetStatus())

}

type User struct {
	UserName string
	Email    string
	Age      int
	IsActive bool
}

func (u User) GetStatus() bool {
	return u.IsActive
}

func (u User) SetStatus(val bool) bool {
	return val
}
