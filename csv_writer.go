package main

import (
	"encoding/csv"
	"os"
)

func main() {
	write := csv.NewWriter(os.Stdout)
	_ = write.Write([]string{"Muhammad", "Al-fatih", "Ritonga"})
	_ = write.Write([]string{"Syilfa", "Salsabila", "Siregar"})
	_ = write.Write([]string{"Syakira", "Shabrina", "Ritonga"})

	write.Flush()
}
