package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	age := r.URL.Query().Get("age")
	fmt.Printf(age)
	fmt.Fprintf(w, "Hello, You have requested the path : %s\n", r.URL.Path)
}

func main() {
	fmt.Printf("Hello world")
	http.HandleFunc("/", home)

	// serving static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)

}
