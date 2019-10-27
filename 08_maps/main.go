package main

import "fmt"

func main() {
	// // Define map
	// me := make(map[string]string)
	//
	// // Asign map
	// me["name"] = "Faisal"
	// me["email"] = "faisal@gmail.com"
	// me["age"] = "34"

	// Declare map and assign
	me := map[string]string{
		"name":  "Faisal",
		"email": "faisal@gmail.com",
		"age":   "56",
	}

	fmt.Println(me["name"])
	delete(me, "age")
	fmt.Println(me)

}
