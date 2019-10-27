package main

import "fmt"

func main() {
	x := 9
	y := 9

	// If else
	if x <= y {
		fmt.Printf("%d is less than or equal%d", x, y)
	} else {
		fmt.Printf("%d is less than %d", y, x)
	}

	// else if
	color := "blue"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")
	}

	// switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not red or blue")
	}

}
