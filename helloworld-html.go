package main

import (
	"html/template"
	"net/http"
)

// simulate a user
type User struct {
	Name string
}

func main() {
	http.HandleFunc(
		"/", // pattern
		func(w http.ResponseWriter, r *http.Request) { // handler
			user := User{Name: "kenny"}
			tpl,
				// _ := template.New("").Parse(
				// `<html>
				// 	<head>
				// 	<title>Page - by Kenny</title>
				// 	</head>
				// 	<body>
				// 	<h1>Hello, {{.Name}}!</h1>
				// 	</body>

				// </html>`)
				_ := template.ParseFiles("templates/hello.html")
			tpl.Execute(w, &user)
		})
	http.ListenAndServe(":8888", nil)
}
