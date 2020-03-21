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

// Post struct
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
			u1 := User{Name: "Kenny"}
			u2 := User{Name: "Bob"}

			posts := []Post{
				Post{User: u1, Body: "Beautiful day in NY!"},
				Post{User: u2, Body: "Beautiful Day in SF!"},
			}

			vm := HelloViewModel{Title: "Home Page", User: u1, Posts: posts} // instantiate VM
			tpl, _ := template.ParseFiles("templates/hello3.html")           // get the template
			tpl.Execute(w, &vm)                                              // use the template
		})

	// func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8888", nil)
}
