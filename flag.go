package main

import (
	"flag"
	"fmt"
)

const (
	expectedUsername = "admin"
	expectedPassword = "password"
)

func Login(username, password string) bool {
	if username == expectedUsername && password == expectedPassword {
		return true
	} else {
		return false
	}
}

func main() {
	username := flag.String("u", "", "your username")
	password := flag.String("p", "", "your password")

	flag.Parse()
	// command prompt
	// go run flag.go -u="admin" -p="password"

	ok := Login(*username, *password)
	if !ok {
		fmt.Println("Invalid username or password.")
	} else {
		fmt.Println("Login successful!")
	}
}
