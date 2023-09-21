package main

import "fmt"

func main() {
    var sisiAtas, sisiBawah, tinggi float64

    fmt.Print("Masukkan panjang sisi atas trapesium: ")
    fmt.Scan(&sisiAtas)

    fmt.Print("Masukkan panjang sisi bawah trapesium: ")
    fmt.Scan(&sisiBawah)

    fmt.Print("Masukkan tinggi trapesium: ")
    fmt.Scan(&tinggi)

    luas := hitungLuasTrapesium(sisiAtas, sisiBawah, tinggi)

    fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}

func hitungLuasTrapesium(sisiAtas, sisiBawah, tinggi float64) float64 {
    return 0.5 * (sisiAtas + sisiBawah) * tinggi
}
