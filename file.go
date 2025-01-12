package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0777)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	// err := createNewFile("sample.txt", "this is sample for create new file")
	// if err != nil {
	// 	fmt.Println("Error", err.Error())
	// } else {
	// 	fmt.Println("Succes create file.")
	// }

	// result, err := readFile("sample.txt")
	// if err != nil {
	// 	fmt.Println("Error", err.Error())
	// } else {
	// 	fmt.Println(result)
	// }

	err := addToFile("sample.txt", "\nthis is new line for sample.txt")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Succes add to file.")
	}
}
