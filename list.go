package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Data 1")
	data.PushBack("Data 2")
	data.PushBack("Data 3")
	data.PushBack("Data 4")
	data.PushBack("Data 5")

	head := data.Front()    // mengambil data urutan pertama dari list
	fmt.Println(head.Value) // Data 1

	// mengambil data selanjutnya
	fmt.Println(head.Next().Value) // Data 2

	tail := data.Back() // mengambil data urutan terakhir dari list
	fmt.Println(tail.Value)

	// fmt.Println(tail.Next().Value) // panic: error karna tidak ada data selanjutnya

	// looping data list
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
