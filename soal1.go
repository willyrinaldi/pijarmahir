package main

import "fmt"

func main() {
    // Buat dua array contoh
    array1 := []string{"Alice", "Bob", "Charlie"}
    array2 := []string{"Bob", "David", "Eve"}

    // Buat peta untuk melacak nama yang sudah ada
    existingNames := make(map[string]bool)

    // Inisialisasi array gabungan
    var combinedArray []string

    // Gabungkan data dari array pertama
    for _, name := range array1 {
        if !existingNames[name] {
            combinedArray = append(combinedArray, name)
            existingNames[name] = true
        }
    }

    // Gabungkan data dari array kedua
    for _, name := range array2 {
        if !existingNames[name] {
            combinedArray = append(combinedArray, name)
            existingNames[name] = true
        }
    }

    // Cetak hasil penggabungan
    fmt.Println("Array yang digabungkan:", combinedArray)
}
