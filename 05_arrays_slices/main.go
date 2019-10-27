package main

import "fmt"

func main() {
	// // Arrays
	// var fruitArr [3]string
	//
	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"
	// fruitArr[2] = "Orange"

	// // Declare and Assign
	// fruitArr := [3]string{"Apple", "Orange", "Mango"}
	//
	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	// Slice
	fruitSlice := []string{"Apple", "Orange", "Mango", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
