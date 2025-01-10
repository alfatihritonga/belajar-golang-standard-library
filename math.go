package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(2.20))            // membulatkan keatas
	fmt.Println(math.Floor(2.80))           // membulatkan kebawah
	fmt.Println(math.Round(2.45))           // membulatkan ke yang terdekat
	fmt.Println(math.Max(10.5, 10.2192091)) // mencari nilai terbesar
	fmt.Println(math.Min(10.5, 10.2192091)) // mencari nilai terkecil
}
