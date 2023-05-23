package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name string
	Id   int
}

func main() {
	u1 := user{Name: "Renuka sai", Id: 19}

	b, err := json.Marshal(u1)
	if err != nil {
		log.Fatalln("JSON did not marshal - and the error is:", err)
	}
	fmt.Println(string(b))

}
