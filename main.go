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
		Schema:   &ReservationSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	/*reservations := handler.New(&handler.Config{
		Schema:   &ReservationsSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	reservation := handler.New(&handler.Config{
		Schema:   &ReservationSchema,
		Pretty:   true,
		GraphiQL: true,
	})*/

	http.Handle("/reserve", reserve)
	//http.Handle("/reservations", reservations)
	//http.Handle("/reservation", reservation)

	http.HandleFunc("/hello", helloWorld)
	fmt.Printf("listening...")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	//log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), reservations))
	//log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), reservation))



}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!!")
}

func myPrint(string string) {
	fmt.Printf(string)
}
