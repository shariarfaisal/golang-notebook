package main

/*
strings.Compare("a","b") // -1 or 0 or 1
strings.Contains("faisal","ais") // true or false
strings.ContainsAny("faisal","i") // true or false
strings.Count("ieiou","i") // 2
strings.EqualFold("go","Go") // true
strings.Fields("i love you") // []string{"i","love","you"}
strings.HasPrefix("Faisal","Fa") // true
strings.HasSuffix("Faisal","al") // true
strings.Index("chicken", "ken") // 4 , -1
strings.IndexAny("chicken", "aeiouy") // 2
strings.Join([]string{"i","love","you"}, ", ")  // "i, love, you"
strings.LastIndex("go gopher", "go") // 3
strings.LastIndexAny("go gopher", "go") //  4
strings.Repeat("na", 3) // nanana
strings.Replace("oink oink oink", "k", "ky", 2)
strings.ReplaceAll("oink oink oink", "oink", "moo")
strings.Split("a,b,c", ",")
strings.SplitAfter("a,b,c", ",")
strings.SplitAfterN("a,b,c,d,d,e,e", ",", 2)
strings.SplitN("a,b,c", ",", 0)
strings.Title("her royal highness") // Her Royal Highness
strings.ToLower("Gopher") //gopher
strings.ToTitle("her royal highness") //HER ROYAL HIGHNESS
strings.ToUpper("Gopher")
strings.Trim("¡¡¡Hello, Gophers!!!", "!¡") //Hello, Gophers
strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡")
*/

func main() {

	// fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))

	// fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
	// 	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	// }))

	// fmt.Println(strings.Compare("a", "b"))
	// fmt.Println(strings.Compare("a", "a"))
	// fmt.Println(strings.Compare("b", "a"))

	// fmt.Println(strings.Contains("faisal", "sal"))
	// fmt.Println(strings.Contains("faisal", "faisal"))
	// fmt.Println(strings.Contains("faisal", "hey"))
	// fmt.Println(strings.Contains("", ""))

	// fmt.Println(strings.ContainsAny("team", "i"))
	// fmt.Println(strings.ContainsAny("fail", "ui"))
	// fmt.Println(strings.ContainsAny("ure", "eru"))
	// fmt.Println(strings.ContainsAny("failure", "fli"))
	// fmt.Println(strings.ContainsAny("foo", ""))
	// fmt.Println(strings.ContainsAny("", ""))

	// fmt.Println(strings.Count("chkleedkeese", "e"))
	// fmt.Println(strings.Count("five", "i")) // before & after each rune

	// fmt.Println(strings.EqualFold("go", "Go"))

	// x := strings.Fields("hello world")
	// fmt.Println(x[1])
	// fmt.Println(len(x))
	// fmt.Printf("%T", x)
	// fmt.Printf("Fields are: %q", strings.Fields("foo bar  baz"))

	// fmt.Println(strings.HasPrefix("Gopher", "Go"))
	// fmt.Println(strings.HasPrefix("Gopher", "C"))
	// fmt.Println(strings.HasPrefix("Gopher", ""))

	// fmt.Println(strings.HasSuffix("Gopher", "er")) //true
	// fmt.Println(strings.HasSuffix("Gopher", "Er")) //false
	// fmt.Println(strings.HasSuffix("Gopher", "C"))  //false
	// fmt.Println(strings.HasSuffix("Gopher", ""))   //true

	// fmt.Println(strings.Index("chicken", "ken"))
	// fmt.Println(strings.Index("chicken", "dmr"))

	// fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	// fmt.Println(strings.IndexAny("crwth", "aeiouy"))

	// s := []string{"foo", "bar", "baz"}
	// fmt.Println(strings.Join(s, ", "))
	// fmt.Println(strings.Join(s, "/ "))

	// fmt.Println(strings.Index("go gopher", "go"))
	// fmt.Println(strings.LastIndex("go gopher", "go"))
	// fmt.Println(strings.LastIndex("go gopher", "rodent"))

	// fmt.Println(strings.LastIndexAny("go gopher", "go"))
	// fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	// fmt.Println(strings.LastIndexAny("go gopher", "fail"))

	// fmt.Println("ba" + strings.Repeat("na", 3))

	// fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	// fmt.Println(strings.Replace("oink oink oink", "oink", "moo", 1))

	// fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	// fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	// fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	// fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	// fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))

	// fmt.Printf("%q\n", strings.SplitAfterN("a,b,c,d,d,e,e", ",", 2))

	// fmt.Printf("%q\n", strings.SplitN("a,b,c,d,f,g,h,", ",", 4))
	// z := strings.SplitN("a,b,c", ",", 0)
	// fmt.Printf("%q (nil = %v)\n", z, z == nil)

	// Compare this example to the ToTitle example.
	// fmt.Println(strings.Title("her royal highness"))
	// fmt.Println(strings.Title("loud noises"))
	// fmt.Println(strings.Title("хлеб"))

	// fmt.Println(strings.ToLower("Gopher"))

	// Compare this example to the Title example.
	// fmt.Println(strings.ToTitle("her royal highness"))
	// fmt.Println(strings.ToTitle("loud noises"))
	// fmt.Println(strings.ToTitle("хлеб"))

	// fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))

	// fmt.Println(strings.ToUpper("Gopher"))

	// fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))

}
