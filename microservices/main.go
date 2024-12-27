package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "das ist die homepage")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "das ist die homepage")
}
func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "das ist die homepage")
}
func Services(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "das ist die homepage")
}

func main() {
	const portN string = ":8000"
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)
	http.HandleFunc("/services", Services)
	fmt.Println("Spinning up")
	http.ListenAndServe(portN, nil)

}
