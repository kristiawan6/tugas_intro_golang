package main

import "fmt"

func Grading(average float64) string {
	if average >= 90 && average <= 100 {
		return "Grade = A"
	} else if average >= 80 && average <= 89 {
		return "Grade = B"
	} else if average >= 70 && average <= 79 {
		return "Grade = C"
	} else if average >= 60 && average <= 69 {
		return "Grade = D"
	} else {
		return "Grade = E"
	}
}

func CalculateAverage(mtk, ipa, bahasaIndonesia, bahasaInggris int) float64 {
	sum := float64(mtk + ipa + bahasaIndonesia + bahasaInggris)/4
	return sum
}
func main() {
	/*
		Terms & Condition
		90 - 100 = A
		80 - 89 = B
		70 - 79 = C
		60 - 69 = D
		0 - 59 = E
	*/

	mtk := 80
	bahasaIndonesia := 90
	bahasaInggris := 89
	ipa := 69

	if mtk == 0 || bahasaIndonesia == 0 || bahasaInggris == 0 || ipa == 0 {
		fmt.Println("Semua nilai harus diisi.")
		return
	}

	average := CalculateAverage(80,90,89,69)
	fmt.Println("Rata - Rata =", average)

	grade := Grading(average)
	fmt.Println(grade)
}
