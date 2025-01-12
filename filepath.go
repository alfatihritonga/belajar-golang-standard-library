package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("belajar-golang-standard-library/path.go"))
	fmt.Println(filepath.Base("Belajar/GOLANG/belajar-golang-standard-library/path.go"))
	fmt.Println(filepath.Ext("Belajar/GOLANG/belajar-golang-standard-library/path.go"))
	fmt.Println(filepath.IsAbs("Belajar/GOLANG/belajar-golang-standard-library/path.go"))
	fmt.Println(filepath.IsLocal("Belajar/GOLANG/belajar-golang-standard-library/path.go"))
	fmt.Println(filepath.Join("Belajar", "GOLANG", "belajar-golang-standard-library", "path.go"))
}
