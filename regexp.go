package main

import (
	"fmt"
	"regexp"
)

type Valid func(string) bool

func isValidPhoneNumber(phone string) bool {
	pattern := `^\+628\d{9,13}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(phone)
}

func checkPhoneNumber(phone string, valid Valid) string {
	v := valid(phone)
	if !v {
		return "Nomor telepon tidak valid"
	}
	return "Nomor telepon valid"
}

func main() {
	fmt.Println(checkPhoneNumber("+62812345678", isValidPhoneNumber))       // nomor kurang dari 9 angka
	fmt.Println(checkPhoneNumber("+62812345678901234", isValidPhoneNumber)) // nomor lebih dari 13 angka
	fmt.Println(checkPhoneNumber("+628123456789012", isValidPhoneNumber))
}
