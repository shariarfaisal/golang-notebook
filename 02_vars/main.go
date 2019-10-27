package main

import "fmt"

func main() {
	// Main types
	// string
	// bool
	// int
	// int int8 int16 int64
	// unit unit8 unit16 unit32 unitptr
	// byte - alias for unit8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// variables

	const name string = "Faisal"
	var age float32 = 21.0
	isCool := true
	isCool = false

	city, village := "Feni", "Jitpur"

	fmt.Println(name, age)
	fmt.Println(city, village)

	fmt.Printf("%T ", name) // check data type with %T
	fmt.Printf("%T", isCool)
	fmt.Printf("%T", age)
}
