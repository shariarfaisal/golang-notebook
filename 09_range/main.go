package main

import "fmt"

func main() {

	// ids := []int{33, 45, 66, 22, 75, 22, 34, 54, 32}
	//
	// // Loop through ids
	// for i, id := range ids {
	// 	fmt.Printf("%d - ID %d", i, id)
	// }

	me := map[string]string{
		"name":  "Faisal",
		"email": "faisal@gmail.com",
		"age":   "56",
	}

	for k, v := range me {
		fmt.Println(k + " : " + v)
	}

	for _, v := range me {
		fmt.Println("value", v)
	}

}
