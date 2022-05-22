package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request()
}

// Createa a basic home page function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the new HomePage!\n")
	fmt.Println("Endpoint Hit: homePage")
}

// Create an about me function
func aboutMe(w http.ResponseWriter, r *http.Request) {
	who := "Keith Lemon"
	fmt.Fprintf(w, "About Me: %s\n", who)
	fmt.Println("Endpoint Hit: aboutMe")
}

// Request function
func request() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutMe)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
