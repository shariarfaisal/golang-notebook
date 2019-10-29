package main

// 2 hour 43 mins left ...
import "fmt"

type address struct {
	Dep   string
	Books []string
	Shift string
}

type person struct {
	address
	Name  string
	Email string
	Phone string
	Gpa   float64
}

func (p person) greeting() {
	fmt.Println("hello " + p.Name)
}

func (p *person) set(value string) {
	p.Name = value
}

func main() {

	// add := address{
	// 	Dep:   "tct",
	// 	Books: []string{"A", "B", "C"},
	// 	Shift: "Morning",
	// }

	faisal := person{
		// address: add, or address{...}
		// Name:  "Faisal",
		Email: "faisal@gmail.com",
		Phone: "018",
		Gpa:   3.75,
	}

	faisal.set("Faisal")

	faisal.Dep = "tct"
	faisal.Books = []string{"A", "B", "C"}
	faisal.Shift = "Morning"

	fmt.Println(faisal.Dep)
	faisal.greeting()

}
