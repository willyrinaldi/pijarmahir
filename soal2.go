package main

import (
	"fmt"
	"math"
)

// Definisikan struct Student
type Student struct {
	Name  string
	Score int
}

// Struct School yang berisi data siswa dan metodenya
type School struct {
	Students []Student
}

// Method untuk menghitung rata-rata skor siswa
func (s School) CalculateAverageScore() float64 {
	totalScore := 0
	for _, student := range s.Students {
		totalScore += student.Score
	}
	averageScore := float64(totalScore) / float64(len(s.Students))
	return averageScore
}

// Method untuk mencari nilai minimum dan nama siswa dengan skor minimum
func (s School) FindMinScore() (string, int) {
	minScore := math.MaxInt
	minStudentName := ""
	for _, student := range s.Students {
		if student.Score < minScore {
			minScore = student.Score
			minStudentName = student.Name
		}
	}
	return minStudentName, minScore
}

// Method untuk mencari nilai maksimum dan nama siswa dengan skor maksimum
func (s School) FindMaxScore() (string, int) {
	maxScore := math.MinInt
	maxStudentName := ""
	for _, student := range s.Students {
		if student.Score > maxScore {
			maxScore = student.Score
			maxStudentName = student.Name
		}
	}
	return maxStudentName, maxScore
}

func main() {
	// Membuat instance School dan mengisi data siswa
	school := School{
		Students: []Student{
			{"Risky", 80},
			{"Kobar", 75},
			{"Ismail", 70},
			{"Umam", 75},
			{"Yopan", 60},
		},
	}

	// Menghitung rata-rata skor
	averageScore := school.CalculateAverageScore()
	fmt.Printf("Average Score: %.2f\n", averageScore)

	// Mencari nilai minimum dan siswa dengan skor minimum
	minStudentName, minScore := school.FindMinScore()
	fmt.Printf("Min Score of Students: %s(%d)\n", minStudentName, minScore)

	// Mencari nilai maksimum dan siswa dengan skor maksimum
	maxStudentName, maxScore := school.FindMaxScore()
	fmt.Printf("Max Score of Students: %s(%d)\n", maxStudentName, maxScore)
}
