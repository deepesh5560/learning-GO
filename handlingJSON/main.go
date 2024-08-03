package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"` // this 3rd parameter will tell which name should be passed in json format (`json:"name"`)
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Pass     string   `json:"Pass"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("hello world")
	encodeJson()
	decodeJson()
}
func encodeJson() {
	courses := []course{
		{"Go Programming", 100, "Udemy", "Free", []string{"Go", "Programming"}},
		{"Python Programming", 150, "Udemy", "Free", nil},
	}

	result, _ := json.MarshalIndent(courses, "", "\t") // this will convert data in JSON but throw data in bits but we can use string() to see the JSON

	fmt.Printf("%s/n", result)
}

func decodeJson() {
	jsonData := []byte(`[
        {"name":"Go Programming","price":100,"platform":"Udemy","pass":"Free","tags":["Go","Programming"]},
        {"name":"Python Programming","price":150,"platform":"Udemy","pass":"Free","tags":null}
    ]`)

	var courses []course
	err := json.Unmarshal(jsonData, &courses) //it will decode the data which we will recieve
	if err != nil {
		fmt.Println("Error unmarshaling JSON:1", err)
		return
	}
	fmt.Println(courses)

	var res []map[string]interface{}
	errors := json.Unmarshal(jsonData, &res)
	if errors != nil {
		fmt.Println("Error unmarshaling JSON:2", err)
		return
	}

	fmt.Println(res)
}
