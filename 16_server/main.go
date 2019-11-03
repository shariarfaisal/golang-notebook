package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type user struct {
	Name     string
	Email    string
	Password string
}

type product struct {
	Title  string
	Body   string
	Price  int64
	Author user
}

var data = []product{
	product{
		Title: "this is product title one",
		Body:  "there will be product description",
		Price: 100,
		Author: user{
			Name:     "Faisal",
			Email:    "faisal@gmail.com",
			Password: "faisal",
		},
	},
	product{
		Title: "this is product title two",
		Body:  "there will be product description",
		Price: 200,
		Author: user{
			Name:     "Farhad",
			Email:    "farhad@gmail.com",
			Password: "farhad",
		},
	},
}

func main() {
	// err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", data)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)

}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", data)
	if err != nil {
		log.Fatalln("template didn't execute:", err)
	}
}
