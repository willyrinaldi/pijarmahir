package main

import "fmt"

// Function untuk menemukan nilai maksimum dan minimum dari array
func findMaxMin(numbers [6]int) (int, int) {
	// Inisialisasi nilai awal maksimum dan minimum
	max := numbers[0]
	min := numbers[0]

	// Loop untuk mencari nilai maksimum dan minimum
	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func main() {
	var numbers [6]int

	// Mengambil input dari pengguna
	fmt.Println("Enter 6 numbers:")
	for i := 0; i < 6; i++ {
		fmt.Print("Input: ")
		fmt.Scan(&numbers[i])
	}

	// Memanggil fungsi findMaxMin
	max, min := findMaxMin(numbers)

	// Menampilkan hasil
	fmt.Printf("%d is the maximum number\n", max)
	fmt.Printf("%d is the minimum number\n", min)
}
