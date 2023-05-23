package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type user struct {
	Name string
	Id   int
}

func toJSON(a interface{}) ([]byte, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return []byte{}, errors.New(fmt.Sprintf("There was an error in to JSON %v ", err))
	}
	return b, nil
}
func main() {
	u1 := user{Name: "Renuka sai", Id: 19}
	b, err := toJSON(u1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(b))
}
