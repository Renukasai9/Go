package main

import "fmt"

// printing values using Sprintf
var x = 42
var y = "James Bond"
var z = true

func main() {
	res := fmt.Sprintf("%v,%v,%v", x, y, z)
	fmt.Println(res)
}
