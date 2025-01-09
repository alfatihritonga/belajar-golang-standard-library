package main

import (
	"errors"
	"fmt"
)

var (
	FatalError  = errors.New("fatal error")
	SyntaxError = errors.New("syntax error")
)

func getError(err string) error {
	if err == "fatal" {
		return FatalError
	}

	if err == "syntax" {
		return SyntaxError
	}

	return nil
}

func main() {
	err := getError("syntax") // parameter => fatal, syntax
	if err != nil {
		if errors.Is(err, FatalError) {
			fmt.Print(FatalError, ": this is fatal error.")
		} else if errors.Is(err, SyntaxError) {
			fmt.Print(SyntaxError, ": this is syntax error.")
		} else {
			fmt.Print("unknown error: this is unknown error.")
		}
	} else {
		fmt.Print("no error.")
	}
}
