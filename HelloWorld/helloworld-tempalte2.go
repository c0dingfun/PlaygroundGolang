package main

import (
	// template implements data-driven templates for generating textual output.
	"html/template"
	"io/ioutil"

	// http provides HTTP client and server implementations.
	"net/http"
	"os"
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

// PopulateTemplates function
func PopulateTemplates() map[string]*template.Template {
	const base = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(base + "/_base.html"))
	dir, err := os.Open(base + "/content")

	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {
		func() {
			f, err := os.Open(base + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + fi.Name() + "'")
			}

			defer f.Close()
			content, err := ioutil.ReadAll(f)

			if err != nil {
				panic("Failed t read content from file '" + fi.Name() + "'")
			}

			tmpl := template.Must(layout.Clone())
			_, err = tmpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "'")
			}
			result[fi.Name()] = tmpl
		}() // self invoke
	}
	return result
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
			// tpl, _ := template.ParseFiles("templates/hello3.html")           // get the template
			// tpl.Execute(w, &vm)                                              // use the template
			templates := PopulateTemplates()
			templates["hello4.html"].Execute(w, &vm)
		})

	// func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8888", nil)
}
