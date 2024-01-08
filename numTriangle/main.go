package main

import "fmt"

func main() {
	for i := 5; i >= 1; i-- {
		row := ""
		for j := 1; j <= i; j++ {
			row += fmt.Sprint(j)
		}
		fmt.Println(row)
	}
}

