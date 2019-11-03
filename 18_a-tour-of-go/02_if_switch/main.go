package main

func main() {
	/*
	   Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
	*/

	/* Basic syntax of if */
	// x := -1
	// if x < 0 {
	// 	fmt.Println("x is less than 0")
	// }

	/* Short syntax of if */
	// if x := -1; x < 0 {
	// 	fmt.Println("X is less than 0")
	// }

	// if x<0{
	//   ...
	// }else {
	//   ...
	// }

	/* Switch case syntax */
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("OS X.")
	// case "linux":
	// 	fmt.Println("Linux.")
	// default:
	// 	fmt.Printf("%s.\n", os)
	// }

	/* Switch with no condition */

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("Good morning!", t)
	// case t.Hour() < 17:
	// 	fmt.Println("Good afternoon.", t)
	// default:
	// 	fmt.Println("Good evening.", t)
	// }

	// defer fmt.Println("world")
	// fmt.Println("hello")

	/* Last in first out */
	// fmt.Println("counting")
	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }
	// fmt.Println("done")

}
