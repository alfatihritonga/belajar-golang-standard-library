package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (a PersonSlice) Len() int {
	return len(a)
}

func (a PersonSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a PersonSlice) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	persons := []Person{
		{Name: "Fatih", Age: 22},
		{Name: "Elma", Age: 25},
		{Name: "Syakira", Age: 12},
		{Name: "Syilfa", Age: 22},
	}

	sort.Sort(PersonSlice(persons))
	fmt.Println(persons)
}
