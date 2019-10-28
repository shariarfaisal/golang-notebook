package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type fci struct {
	Title  string
	Person map[string]string
}

func main() {

	cst := fci{
		Title: "Computer Science & Technology",
		Person: map[string]string{
			"ashek":  "ashek@gmail.com",
			"tahsin": "tahsin@gmail.com",
			"emon":   "emon@gmail.com",
		},
	}
	dnt := fci{
		Title: "Data Networking and Telecommunication Technology",
		Person: map[string]string{
			"shishir": "shishir@gmail.com",
			"munna":   "munna@gmail.com",
			"parvase": "parvase@gmail.com",
		},
	}
	tct := fci{
		Title: "Telecommunication Technology",
		Person: map[string]string{
			"faisal": "faisal@gmail.com",
			"hridoy": "hridoy@gmail.com",
			"mezbah": "mezbah@gmail.com",
		},
	}

	// data := struct {
	// 	Cst fci
	// 	Dnt fci
	// 	Tct fci
	// }{
	// 	cst,
	// 	dnt,
	// 	tct,
	// }

	data := map[string]fci{
		"cst": cst,
		"dnt": dnt,
		"tct": tct,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
