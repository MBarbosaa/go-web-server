package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parsefrom() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	// Instance variable of http server
	var fileServer http.Handler = http.FileServer(http.Dir("./static"))
	var serverPort string = ":8080"

	// Handles the root dir to the file server
	http.Handle("/", fileServer)

	// Handle the specified roots to their respective functions
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the server itself at the specified port handlig the possible errors
	fmt.Println("Starting server at port", serverPort)
	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatal(err)
	}

}
