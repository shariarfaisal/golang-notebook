package main

import "fmt"

func main() {

	a := 3
	b := &a

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&a)

	*b = 5
	fmt.Println(a)
}
