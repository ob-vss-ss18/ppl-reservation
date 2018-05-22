package main

import "github.com/graphql-go/graphql"

var (
	ReservationSchema graphql.Schema
	reservationType *graphql.Object

	Reservations  map[int]Reservation
)

type Reservation struct {
	id int
	cId int
	itemId int
}



func InitializeUserDB() {

	Res = funcName()

	reservationType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Reservation",
		Description: "This is a reservation.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The id of the reservation.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					reservation, ok := p.Source.(Reservation);
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
					reservation, ok := p.Source.(Reservation);
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
					reservation, ok := p.Source.(Reservation);
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
					reservationSlice := make([]Reservation, len(Reservations))
					idx := 0
					for _, reservation := range Reservations {
						if(reservation.cId == cId) {
							reservationSlice[idx] = reservation
							idx++
						}
					}
					return reservationSlice, nil
				},
			},
		},
	})

	ReservationSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,})
}


func GetReservation(cId int) Reservation {
	return Reservation{1,1,1}
}