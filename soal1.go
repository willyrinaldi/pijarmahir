package main

import (
	"fmt"
)

// Definisikan struct Car
type Car struct {
	Type    string
	FuelIn  float64 // Bahan bakar dalam liter
}

// Method untuk menghitung perkiraan jarak
func (c Car) CalculateRange() float64 {
	// Efisiensi bahan bakar (1.5 L/mil)
	fuelEfficiency := 1.5
	// Menghitung jarak yang dapat ditempuh
	rangeInMiles := c.FuelIn * fuelEfficiency
	return rangeInMiles
}

func main() {
	// Membuat instance Car
	myCar := Car{
		Type:   "Sedan",
		FuelIn: 10.0, // Isi bahan bakar dalam liter
	}

	// Memanggil method CalculateRange untuk menghitung jarak yang dapat ditempuh
	rangeEstimation := myCar.CalculateRange()
	fmt.Printf("Jarak perkiraan yang dapat ditempuh dengan mobil %s adalah %.2f mil.\n", myCar.Type, rangeEstimation)
}
