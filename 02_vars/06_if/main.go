package main

import "fmt"

type person struct {
	fname string
	lname string
	email string
	dep   string
}

func (p person) greeting() {
	fmt.Println("hello ", p.lname)
}

func main() {

	faisal := person{
		fname: "Sharia Emon",
		lname: "Faisal",
		email: "faisal623@gmail.com",
		dep:   "tct",
	}

	fmt.Println(faisal.fname)
	faisal.greeting()

	// me := map[string]string{
	// 	"fname": "Faisal",
	// 	"email": "faisal@gmail.com",
	// 	"dep": "tct",
	// 	"shift": "1st",
	// }
	//
	// if email, ok := me["email"]; ok {
	// 	fmt.Println(email, ok)
	// }
	//
	// if name, ok := me["nam"]; ok {
	// 	fmt.Println(name, ok)
	// }

}
