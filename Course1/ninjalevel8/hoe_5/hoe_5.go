package main

import (
	"fmt"
	"sort"
)

// using sort to sort based on age and last
type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int { return len(a) }
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type ByLast []user

func (l ByLast) Len() int { return len(l) }
func (l ByLast) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l ByLast) Less(i, j int) bool {
	return l[i].Last < l[j].Last
}

func main() {
	u1 := user{First: "Renuka sai", Last: "M", Age: 21, Sayings: []string{"how", "hii"}}
	u2 := user{First: "Nandini priya", Last: "Masina", Age: 19, Sayings: []string{"byee", "everyone"}}
	users := []user{u1, u2}
	sort.Sort(ByAge(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, k := range v.Sayings {
			fmt.Println(k, "\t")
		}
	}
	fmt.Println()
	sort.Sort(ByLast(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		for _, k := range v.Sayings {
			fmt.Println(k, "\t")
		}
	}
}
