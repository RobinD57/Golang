package main

import(
	"fmt"
	"net/http"
	)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, Golang style!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "As you can see, my passion is front end!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)
}
