package main

import "net/http"

// Lesson learned: between "package" and "import", there must be an empty
// line; otherwise, "go" will complain "http" is undefined.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8888", nil)
}
