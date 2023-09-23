package main

import "fmt"

func countOccurrences(slice []string, target string) int {
	count := 0
	for _, str := range slice {
		if str == target {
			count++
		}
	}
	return count
}

func main() {
	// Buat sebuah slice contoh
	data := []string{"apel", "jeruk", "apel", "mangga", "apel", "pisang"}

	// String yang ingin dihitung
	target := "apel"

	// Hitung berapa kali string tersebut muncul dalam slice
	occurrences := countOccurrences(data, target)

	// Cetak hasil perhitungan
	fmt.Printf("Jumlah kemunculan '%s' dalam slice: %d\n", target, occurrences)
}
