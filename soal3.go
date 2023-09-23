package main

import (
	"fmt"
)

func findUniqueNumbers(input string) []int {
	// Buat peta (map) untuk menghitung jumlah kemunculan masing-masing angka
	countMap := make(map[rune]int)

	// Hitung jumlah kemunculan setiap angka
	for _, char := range input {
		countMap[char]++
	}

	// Buat slice untuk menyimpan angka yang hanya muncul satu kali
	var uniqueNumbers []int

	// Cari angka yang hanya muncul satu kali dan tambahkan ke slice uniqueNumbers
	for _, char := range input {
		if countMap[char] == 1 {
			uniqueNumbers = append(uniqueNumbers, int(char-'0'))
		}
	}

	return uniqueNumbers
}

func main() {
	// Tes kasus
	input1 := "76523752"
	output1 := findUniqueNumbers(input1)
	fmt.Printf("Input: %s, Output: %v\n", input1, output1)

	input2 := "1122"
	output2 := findUniqueNumbers(input2)
	fmt.Printf("Input: %s, Output: %v\n", input2, output2)
}
