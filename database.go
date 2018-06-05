package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func connect(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS reservations (
      id       SERIAL,
      cId INTEGER,
	  itemId INTEGER,
	  date_from VARCHAR(32),
	  date_to VARCHAR(32)
    );
  `)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getReservation(db *sql.DB, id int) (Reservation, error) {

	var reservation Reservation

	err = db.QueryRow(`SELECT * FROM reservations WHERE id = $1`, id).Scan(&reservation)
	if err == sql.ErrNoRows {
		log.Fatal("No reservation found")
	}
	if err != nil {
		log.Fatal(err)
	}

	return reservation, nil
}

func getReservations(db *sql.DB, id int) ([]Reservation, error) {
	rows, err := db.Query(
		`SELECT * FROM reservations WHERE cId = $1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	reservations := make([]Reservation, 0, 5)

	for rows.Next() {
		u := Reservation{}

		err = rows.Scan(&u.id, &u.cId, &u.itemId, &u.date_from, &u.date_to)

		if err != nil {
			return nil, err
		}

		reservations = append(reservations, u)
	}

	return reservations, nil
}

func setReservation(db *sql.DB, cId int, itemId int, date_from string, date_to string) (bool, error) {

	var id int
	err := db.QueryRow(`INSERT INTO reservations(cId, itemId, date_from, date_to)
	VALUES($1, $2, $3, $4) RETURNING id`, cId, itemId, date_from, date_to).Scan(&id)

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
