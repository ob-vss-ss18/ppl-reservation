package main

import "github.com/graphql-go/graphql"

var (
	ReservationSchema graphql.Schema
	reservationType   *graphql.Object

	Reservations map[int]Reservation
)

type Reservation struct {
	id     int
	cId    int
	itemId int
}

func initReservations() {

	Res1 := Reservation{
		id:     1,
		cId:    1,
		itemId: 1,
	}

	Res2 := Reservation{
		id:     2,
		cId:    1,
		itemId: 2,
	}

	Res3 := Reservation{
		id:     3,
		cId:    2,
		itemId: 3,
	}

	Reservations = map[int]Reservation{
		0: Res1,
		1: Res2,
		2: Res3,
	}

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
					var reservationSlice []Reservation
					idx := 0
					for _, reservation := range Reservations {
						if (reservation.cId == cId) {
							reservationSlice = append(reservationSlice, reservation)
							idx++
						}
					}
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
					idx := 0
					//toDo get information from STOCK check if item is already reserved
					for _, reservation := range Reservations {
						if (reservation.itemId == itemId) {
							return false, nil
							idx++
						}
					}
					Reservations[len(Reservations)] = Reservation{len(Reservations), cId, itemId}
					return true, nil
				},
			},
		},
	})

	ReservationSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,})
}

func GetReservation(cId int) Reservation {
	return Reservation{1, 1, 1}
}
