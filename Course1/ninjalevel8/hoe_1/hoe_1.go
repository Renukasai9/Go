package main

import (
	"encoding/json"
	"fmt"
	"os"
)
//USING marshal
type user struct {
	Name string
	Id   int
}

func main() {
	u1 := user{Name: "Renuka sai", Id: 19}
	u2 := user{Name: "Nandini priya", Id: 18}
	users := []user{u1, u2}
	fmt.Println(users)
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error", err)
	}
	os.Stdout.Write(b)

}
