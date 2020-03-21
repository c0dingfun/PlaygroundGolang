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

type Post struct {
	User User
	Body string
}

// HelloViewModel struct
type HelloViewModel struct {
	Title string
	User  User
	Posts []Post
}

func main() {
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc(
		"/", // pattern
		func(w http.ResponseWriter, r *http.Request) { // handler
			user := User{Name: "kenny"}                            // instantiate User
			vm := HelloViewModel{Title: "Home Page", User: user}   // instantiate VM
			tpl, _ := template.ParseFiles("templates/hello2.html") // get the template
			tpl.Execute(w, &vm)                                    // use the template
		})

	// func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8888", nil)
}
