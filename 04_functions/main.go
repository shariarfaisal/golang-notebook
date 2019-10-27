package main

import "fmt"

// func name(argument argType) returnType{
//   return "Something"
// }

func greeting(name string) string {
	return "Hello " + name
}

func getSum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Sharia emon"))
	fmt.Println(getSum(3, 4))
}
