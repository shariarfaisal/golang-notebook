package main

/** types of value ...
Main types
string
bool
int
int int8 int16 int64
unit unit8 unit16 unit32 unitptr
byte - alias for unit8
rune - alias for int32
float32 float64
complex64 complex128
*/

/*
.Variable declaration...
	var foo int
	var foo int = 34
	foo := 45

.Can't redeclare variables, but can shadow them
.All variables must be used
.Visibility
	.. lower case first letter for package scope
	.. upper case first letter to export
	.. no private scope
*/

/*

SUMMARY
.. Naming conventions
	..Pascal or camelCase
		. Capitalize acronyms (HTTP,URL)
	.. As short as reasonable
		. longer names for longer lives

*/

/*
.Type conversions
	.destinationType(variable)
	.use strconv package for strings
	like..
	var i float32 = 34.3
	var j int = Int(i)
	var s string = strconv.Itoa(i) // import strconv package..
*/

func main() {

	// ways of declare variables....
	// const hello string = "Hello world"
	// var hello string = "Hello world"
	// 		hello := "Hello world"

	// const (
	// 	google string = "http://www.google.com"
	// 	facebook string = "http://www.facebook.com/x_fa1sal"
	// )
	//
	// var (
	// 	fname string = "Sharia Emon"
	// 	lname string = "Faisal"
	// )

	// fmt.Println(fname)

}
