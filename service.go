package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Printf("listening...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!!")
}
