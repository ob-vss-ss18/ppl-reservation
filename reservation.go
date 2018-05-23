package main

import (
	"github.com/graphql-go/graphql"
	"database/sql"
	"os"
	"log"
)

var (
	ReservationSchema graphql.Schema
	reservationType   *graphql.Object

	Reservations map[int]Reservation

	db  *sql.DB
	err error
)

type Reservation struct {
	id     int
	cId    int
	itemId int
}

func initDatabase() {
	url, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatalln("$DATABASE_URL is required")
	}

	db, err = connect(url)
	if err != nil {
		log.Fatalf("Connection error: %s", err.Error())
	}
}

func initGraphQl() {

	initDatabase()

	/**
	Create a graphql object
	 */
	reservationType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Reservation",
		Description: "This is a reservation.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the reservation.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					reservation, ok := p.Source.(Reservation)
					if ok {
						return reservation.id, nil
					}
					return nil, nil
				},
			},
			"cId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of a costumer!!! ",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					reservation, ok := p.Source.(Reservation)
					if ok {
						return reservation.cId, nil
					}
					return nil, nil
				},
			},
			"itemId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of a item!!! ",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					reservation, ok := p.Source.(Reservation)
					if ok {
						return reservation.itemId, nil
					}
					return nil, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"reservations": &graphql.Field{
				Type: graphql.NewList(reservationType),
				Args: graphql.FieldConfigArgument{
					"cId": &graphql.ArgumentConfig{
						Description: "id of the customer",
						Type:        graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					cId := p.Args["cId"].(int)

					var reservationSlice []Reservation
					reservationSlice, err = getReservations(db, cId)

					return reservationSlice, nil
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"reserve": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"cId": &graphql.ArgumentConfig{
						Description: "id of the customer",
						Type:        graphql.NewNonNull(graphql.Int),
					},
					"itemId": &graphql.ArgumentConfig{
						Description: "id of the item",
						Type:        graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					cId := p.Args["cId"].(int)
					itemId := p.Args["itemId"].(int)

					//toDo get information from STOCK check if item is already reserved

					var reserved bool
					reserved, err = setReservation(db, cId, itemId)

					return reserved, err
				},
			},
		},
	})

	ReservationSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,})
}
