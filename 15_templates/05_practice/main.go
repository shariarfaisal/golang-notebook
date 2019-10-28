package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	const tpl = `
  <!DOCTYPE html>
  <html lang="en" dir="ltr">
    <head>
      <meta charset="utf-8">
      <title></title>
    </head>
    <body>
		{{ $title := .Title  }}
      <h1>{{$title}}</h1>
			<ul>
				{{range .Items}}
					<li>{{.}}</li>
				{{end}}
			</ul>
    </body>
  </html>

  `

	check := func(err error) {
		if err != nil {
			log.Fatalln(err)
		}
	}

	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)

}
