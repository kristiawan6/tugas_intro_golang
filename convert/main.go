package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//String to Int
	// str := "william"

	// num, _ := strconv.Atoi(str)
	// fmt.Println("Tipe data :", reflect.TypeOf(str))
	// fmt.Println("Nilai:", num)

	//Int to String
	// age := 19

	// str := strconv.Itoa(age)
	// fmt.Println("Tipe data :", reflect.TypeOf(str))
	// fmt.Println(str)

	//String to Float
	// str := "William"

	// num, _ := strconv.ParseFloat(str, 64)
	// fmt.Println("Tipe data :", reflect.TypeOf(num))
	// fmt.Println("Nilai:", num)

	//Float to String
		height := 169.75

		str := strconv.FormatFloat(height, 'f', 2, 64) // 'f' untuk format desimal, 2 untuk jumlah angka di belakang koma, 64 untuk bit size
		fmt.Println("Tipe data :", reflect.TypeOf(str))
		fmt.Println("Nilai string:", str)
	
}
