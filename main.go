package main

import (
	"fmt"
	"net/http"
	"github.com/graphql-go/handler"
	"log"
)

func main() {

	initReservations()

	h := handler.New(&handler.Config{
		Schema:   &ReservationSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/reservations",h)
	http.Handle("/reserve",h)
	http.HandleFunc("/", helloWorld)
	fmt.Printf("listening...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!!")
}

func myPrint(string string){
	fmt.Printf(string)
}