package main

import (
	"fmt"
	"net/http"
	"github.com/graphql-go/handler"
	"os"
	"log"
)

func main() {
	h := handler.New(&handler.Config{
		Schema:   &ReservationSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	//toDo change /query to something other!!
	http.Handle("/query",h)
	http.HandleFunc("/", helloWorld)
	fmt.Printf("listening...")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!!")
}

func myPrint(string string){
	fmt.Printf(string)
}