package main

import (
	"fmt"
)

func main() {

	// var a = 4
	// var b int8 = 8
	//
	// fmt.Println(a + int(b)) // b converted int8 to int ...

	// var a = 10 // 1010
	// var b = 3  // 0011
	//
	// fmt.Println(a & b)  // 0010 = 2
	// fmt.Println(a | b)  // 1011 = 11
	// fmt.Println(a ^ b)  // 1001 = 9
	// fmt.Println(a &^ b) // 0100 = 8

	// a := 8              // 2^3
	// fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
	// fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1

	a := 10.2
	b := 3.7

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

}
