package main

import (
	"fmt"
	"strconv"
)

type person struct {
	// fname string
	// lname string
	// city  string
	// age   int

	fname, lname, city string
	age                int
}

func (p person) echo(arg string) string {
	return arg + "My name is " + p.fname + p.lname + " and I am  " + strconv.Itoa(p.age)
}

func (p *person) increaseBirthday() {
	p.age++
}

func (p *person) changeName(fname, lname string) {
	p.fname = fname
	p.lname = lname
}

func main() {
	// Init person using struct ...
	fname, lname, age, city := "Sharia Emon", "Faisal", 21, "Feni"
	faisal := person{fname, lname, city, age}

	faisal.increaseBirthday()
	faisal.changeName("Md ", "Al-Arafat")
	fmt.Println(faisal.echo("Hello!"))
}
