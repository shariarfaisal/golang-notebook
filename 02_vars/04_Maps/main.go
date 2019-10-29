package main

import "fmt"

func main() {

	myMap := map[string]string{
		"fname": "Sharia Emon",
		"lname": "Faisal",
		"roll":  "34",
		"dep":   "tct",
	}

	myMap["shift"] = "morning"

	myMap2 := myMap
	delete(myMap2, "roll")

	fname, fok := myMap["fnam"]
	lname, lok := myMap["lname"]

	roll := myMap2["roll"]

	fmt.Printf("%T %v", roll, roll)
	fmt.Println(fname, fok)
	fmt.Println(lname, lok)

	fmt.Println(myMap)
	fmt.Println(myMap2)
}
