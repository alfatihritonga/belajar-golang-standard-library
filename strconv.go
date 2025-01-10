package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res, err := strconv.Atoi("12345")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%d + %d = %d\n", res, res, res*2)
	}

	resItoa := strconv.Itoa(12345)
	fmt.Println(resItoa)
	split := strings.Split(resItoa, "")
	fmt.Println(split)
}
