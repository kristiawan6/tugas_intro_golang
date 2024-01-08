package main

import (
	"fmt"
	"reflect"
)

func main() {
	//name(String)
	var name string = "William"
	fmt.Println("Tipe data :",reflect.TypeOf(name))
	fmt.Println(name)
	
	//age (int)
	var age int = 19
	fmt.Println("Tipe data :",reflect.TypeOf(age))
	fmt.Println(age)

	var height float64 = 169.75
	fmt.Println("Tipe data :",reflect.TypeOf(height))
	fmt.Println(height)

	//IsMarried (Bool)
	var IsMarried bool = false
	fmt.Println("Tipe data :",reflect.TypeOf(IsMarried))
	fmt.Println(IsMarried)
	
	//InterestInCoding (Bool)
	var InterestInCoding bool = true
	fmt.Println("Tipe data :",reflect.TypeOf(InterestInCoding))
	fmt.Println(InterestInCoding)

}