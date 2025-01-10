package main

import (
	"fmt"
	"strings"
)

func FilterBadWords(w string) string {
	w = strings.ToLower(w)

	if strings.ContainsAny(w, "fuck") {
		w = strings.ReplaceAll(w, "fuck", "****")
	}

	if strings.ContainsAny(w, "fck") {
		w = strings.ReplaceAll(w, "fck", "***")
	}

	return w
}

func main() {
	filterbadword := FilterBadWords("FuCK fCK fuck")
	fmt.Println(filterbadword)
}
