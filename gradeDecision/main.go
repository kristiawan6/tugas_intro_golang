package main

import (
	"fmt"
	"reflect"
)

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

func CalculateAverage(mtk, ipa, bahasaIndonesia, bahasaInggris any) float64 {

	mtkScore := mtk.(int)
	ipaScore := ipa.(int)
	bahasaIndonesiaScore := bahasaIndonesia.(int)
	bahasaInggrisScore := bahasaInggris.(int)

	if mtkScore == 0 || bahasaIndonesiaScore == 0 || bahasaInggrisScore == 0 || ipaScore == 0 {
		return 0
	}

	sum := float64(mtkScore+ipaScore+bahasaIndonesiaScore+bahasaInggrisScore) / 4
	return sum
}

func validateType(value any, targetType string) bool {
	valueType := reflect.TypeOf(value).Kind().String()
	return valueType == targetType
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
	bahasaInggris := "89"
	ipa := 69

	//Validasi tipe data
	isMTKInt := validateType(mtk, "int")
	if !isMTKInt {
		fmt.Println("Nilai MTK harus dalam tipe data angka")
		return
	}
	isbahasaIndonesiaInt := validateType(bahasaIndonesia, "int")
	if !isbahasaIndonesiaInt {
		fmt.Println("Nilai Bahasa Indonesia harus dalam tipe data angka")
		return
	}
	isbahasaInggrisInt := validateType(bahasaInggris, "int")
	if !isbahasaInggrisInt {
		fmt.Println("Nilai bahasa Inggris harus dalam tipe data angka")
		return
	}
	isIpaInt := validateType(ipa, "int")
	if !isIpaInt {
		fmt.Println("Nilai ipa harus dalam tipe data angka")
		return
	}

	average := CalculateAverage(mtk, bahasaIndonesia, bahasaInggris, ipa)
	if average == 0 {
		fmt.Println("Semua nilai harus diisi.")
		return
	}
	fmt.Println("Rata - Rata =", average)

	grade := Grading(average)
	fmt.Println(grade)
}
