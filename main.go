package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorld)
	fmt.Printf("listening...")
	log.Fatal(http.ListenAndServe(":" + os.Getenv("Path"), nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!!")
}
