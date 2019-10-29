package main

import "fmt"

/*
.Slices
  .Backed by array
  .creation styles
    .Slice existing array or slice
    .literal style
    .via make function
      - a:= make([]int,10) // create slice with capacity and length == 10
      - a:= make([]int,10,100) // slice with length == 10 and capacity == 100
  .len function returns length of slice
  .cap function returns length of underlying array
  .append function to add elements to slice
    .May cause expensinve copy operation if underlying array is too small

*/

func main() {
	a := [...]int{1, 2, 3, 4}
	b := a
	a[2] = 8
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
}
