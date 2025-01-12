package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Fatih", "Syilfa", "Syakira", "Elma"}
	values := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Contains(names, "Fatih"))
	fmt.Println(slices.Contains(names, "Eko"))
	fmt.Println(slices.Index(names, "Fatih"))
	fmt.Println(slices.Index(names, "Eko"))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
}
