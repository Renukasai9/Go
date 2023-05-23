package main

import (
	"encoding/json"
	"fmt"
)

// using unmarshal
type user struct {
	Name string
	Id   int
}

func main() {
	a := `[{"Name":"Renuka sai","Id":19},{"Name":"Nandini priya","Id":18}]`
	var users = []byte(a)
	var result []user
	err := json.Unmarshal(users, &result)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("%+v \n", result)
}
