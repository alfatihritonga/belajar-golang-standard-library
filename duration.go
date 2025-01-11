package main

import (
	"fmt"
	"time"
)

func main() {
	milliSecond := time.Millisecond
	second := time.Second
	minute := time.Minute
	hour := time.Hour

	fmt.Println(milliSecond)
	fmt.Println(second)
	fmt.Println(minute)
	fmt.Println(hour)

	fmt.Println(hour + minute + second + milliSecond)
}
