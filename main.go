package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Found", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse(), err %v", err)
		return
	}
	fmt.Fprintf(w, "Post Reqest Successfull")
	name := r.FormValue("name")
	addess := r.FormValue("address")
	fmt.Fprintf(w, "\n name is %s and address is %s", name, addess)
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
