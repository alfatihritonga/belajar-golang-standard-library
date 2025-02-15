package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Muhammad,Al-fatih,Ritonga\n" +
		"Syilfa,Salsabila,Siregar\n" +
		"Syakira,Shabrina,Ritonga"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
