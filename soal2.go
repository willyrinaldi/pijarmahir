package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan sebuah bilangan: ")
    fmt.Scan(&n)

    if isMultipleOf7(n) {
        fmt.Printf("%d adalah kelipatan 7\n", n)
    } else {
        fmt.Printf("%d bukan kelipatan 7\n", n)
    }
}

func isMultipleOf7(n int) bool {
    return n%7 == 0
}
