package main

import "fmt"

func main() {
	const name, age = "Al-fatih", 22
	fmt.Print("Hi, i'm ", name, " and i'm ", age, " years old.\n") // print output tanpa format
	fmt.Printf("Hi, i'm %s and i'm %d years old.\n", name, age)    // print menggunakan format
}
