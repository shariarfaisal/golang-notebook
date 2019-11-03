package main

import "fmt"

func main() {
	i, j := 34, 567
	p := &i

	*p = 12
	p = &j
	*p = 78
	fmt.Println(i, j)
}
