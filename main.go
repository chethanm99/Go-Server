package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST Request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name= %s\n", name)
	fmt.Fprintf(w, "Address =%s\n", address)
}

func wellbeingHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST Request sucessful\n")
	name := r.FormValue("name")
	age := r.FormValue("age")
	health_condition := r.FormValue("health-condition")
	fmt.Fprintf(w, "Name= %s\n", name)
	fmt.Fprintf(w, "Health Condition=%s\n", health_condition)
	fmt.Fprintf(w, "Age= %s\n", age)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/goodbye" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method not found", http.StatusNotFound)
	}
	fmt.Fprintf(w, "good bye thank you for your time, it was great spending with you")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/goodbye", goodbyeHandler)
	http.HandleFunc("/wellbeing", wellbeingHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(": 8080", nil); err != nil {
		log.Fatal(err)
	}
}
