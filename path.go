package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("belajar-golang-standard-library/path.go"))
	fmt.Println(path.Base("Belajar/GOLANG/belajar-golang-standard-library/path.go"))
	fmt.Println(path.Ext("Belajar/GOLANG/belajar-golang-standard-library/path.go"))
	fmt.Println(path.Join("Belajar", "GOLANG", "belajar-golang-standard-library", "path.go"))
}
