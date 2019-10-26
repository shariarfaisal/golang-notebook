package main

import "fmt"

type address struct {
	country  string
	city     string
	district string
}

type people struct {
	address
	fname string
	lname string
	age   int
	dep   string
	drink bool
}

func (pe people) hello() { // method
	fmt.Println(pe.fname, pe.lname, pe.address.country, `hello world `)
}

func (pe people) welcome() {
	fmt.Println(pe.lname, "Welcome")
}

// Polymorphism ....

type human interface {
	hello()
	welcome()
}

func saySomething(h human) {
	h.welcome()
}

func main() {

	// Declare variables
	// var country string
	// country = "Bangladesh"
	// var dst = "Feni"
	// institute := "Feni Computer Institute"
	// village, postOffice, union := "Jitpur", "Motigonj", "Motigonj"
	//
	// book := []string{"Math", "English", "Physics"} //slice
	// bookPrice := map[string]int{                   // map
	// 	"Math":    200,
	// 	"English": 150,
	// 	"Physics": 300,
	// }
	//
	// fmt.Println(institute, country, dst)
	// fmt.Println(village, postOffice, union)
	// fmt.Println(book[0])
	// fmt.Println(bookPrice["Math"])

	me := people{ //struct
		address{
			"Bangladesh",
			"Feni",
			"Feni",
		},
		"Sharia Emon",
		"Faisal",
		21,
		"tct",
		false,
	}

	brother := people{
		address{
			"BD",
			"Feni",
			"Feni",
		},
		"Mezbah",
		"Uddin",
		22,
		"tct",
		false,
	}

	// me.hello()

	saySomething(me)
	saySomething(brother)

}
