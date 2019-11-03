package main

import "fmt"

// type char struct {
// 	A string
// 	B int
// }

func main() {
	// x := char{
	// 	"Faisal",
	// 	34,
	// }
	//
	// y := &x
	// y.A = "Farhad"
	//
	// z := &char{"helo", 4}
	//
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

	// type hello struct {
	// 	str string
	// 	boo bool
	// }

	// x := []struct {
	// 	str string
	// 	boo bool
	// }{
	// 	{str: "hello", boo: true},
	// 	{str: "hey", boo: false},
	// 	{str: "hi", boo: false},
	// 	{str: "lili", boo: true},
	// 	{str: "what the heck", boo: false},
	// 	{str: "lol", boo: true},
	// }

	// x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//
	// y := x[:4]

	// fmt.Println(x)
	// fmt.Println(y)

	// printSlice(y)

	// var s []int
	// fmt.Println(s, len(s), cap(s))

	// i := make([]int, 5, 9) // ([]int,value,capacity)
	// printSlice(i)
	//
	// var s []int
	// printSlice(s)
	//
	// // append works on nil slices.
	// s = append(s, 0)
	// printSlice(s)
	//
	// // The slice grows as needed.
	// s = append(s, 1)
	// printSlice(s)
	//
	// // We can add more than one element at a time.
	// s = append(s, 2, 3, 4)
	// printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d capacity=%d value=%v \n", len(s), cap(s), s)
}
