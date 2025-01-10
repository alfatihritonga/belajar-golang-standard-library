package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (a UserSlice) Len() int {
	return len(a)
}

func (a UserSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a UserSlice) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	users := []User{
		{Name: "Fatih", Age: 22},
		{Name: "Elma", Age: 25},
		{Name: "Syakira", Age: 12},
		{Name: "Syilfa", Age: 22},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
