package main

import (
	"fmt"      // I/O functions. Basically console stuff
	"log"      // Just a logger
	"net/http" // The HTTP server
)

// Handles the hello world route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Checks that the petition is in the correct method
	if r.Method != "GET" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Reply with Hello World!
	fmt.Fprintf(w, "Hello World!")
}

// Handles the form route
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Checks that the petition is in the correct method
	if r.Method != "POST" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Checks if the body can be parsed
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Something went wrong: %v", err)
		return
	}

	// In Go, you can "reply" several times and all will be concatenated at the end
	fmt.Fprintf(w, "Everything went okey. \n")

	// Gets the "name" and "address" values
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s \n", name)
	fmt.Fprintf(w, "Address: %s \n", address)
}

func main() {
	// := is a short declaration. It infers the type, declares the variable and assigns a value
	fileServer := http.FileServer(http.Dir("./static"))

	// Handle the / route with fileServer
	http.Handle("/", fileServer)

	// Handle the routes with a function rather than a handler
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080")

	// First declare and assign the err variable, then check if its different than nil
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
