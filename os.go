package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	// mendapatkan home directory
	userhomedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userhomedir)

	// mendapatkan hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(hostname)
}
