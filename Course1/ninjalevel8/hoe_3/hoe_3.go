package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// using encoder
type user struct {
	Name string
	Id   int
}

func main() {
	u1 := user{Name: "Renuka sai", Id: 19}
	u2 := user{Name: "Nandini priya", Id: 18}
	users := []user{u1, u2}
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("no error")
	}

}
