package main

import (
	// template implements data-driven templates for generating textual output.
	"html/template"

	// http provides HTTP client and server implementations.
	"net/http"
)

// User struct
type User struct {
	Name string
}

func main() {
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc(
		"/", // pattern
		func(w http.ResponseWriter, r *http.Request) { // handler
			user := User{Name: "kenny"}

			//tpl, _ := template.New("").Parse(
			// `<html>
			// 	<head>
			// 	<title>Page - by Kenny</title>
			// 	</head>
			// 	<body>
			// 	<h1>Hello, {{.Name}}!</h1>
			// 	</body>
			// </html>`)

			// use separated html  (tpl = template)
			tpl, _ := template.ParseFiles("templates/hello.html") // get the template from file
			tpl.Execute(w, &user)                                 // use the template
		})

	// func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8888", nil)
}
