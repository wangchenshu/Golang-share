package main

import (
	"fmt"
	"sample/hello"
)

func GetNums() (int, int) {
	return 1, 2
}

func main() {
	helloStr := hello.SayHello()
	fmt.Println(helloStr)

	// var x int = 10
	// fmt.Println(x)

	x, y, z := 10, 3.14, "Walter"
	fmt.Println(x, y, z)

	text1 := "Go語言"
	text2 := "Cool"
	var text3 string

	fmt.Println(text1 + text2) // Go語言Cool
	fmt.Printf("%q\n", text3)  // ""
	fmt.Println(text1 > text2) // true

	num1, num2 := GetNums()
	fmt.Println(num1, num2)
}
