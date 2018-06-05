package main

import (
	"fmt"
	"net/http"
	"github.com/graphql-go/handler"
	"log"
	"os"
)

func main() {

	initGraphQl()

	reserve := handler.New(&handler.Config{
		Schema:   &MutationSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	reservations := handler.New(&handler.Config{
		Schema:   &ReservationsSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	reservation := handler.New(&handler.Config{
		Schema:   &ReservationSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/", reserve)
	http.Handle("/", reservations)
	http.Handle("/", reservation)

	http.HandleFunc("/hello", helloWorld)
	fmt.Printf("listening...")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))


}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!!")
}

func myPrint(string string) {
	fmt.Printf(string)
}
