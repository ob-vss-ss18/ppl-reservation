package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var ()

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
	  itemId INTEGER
    );
  `)
	if err != nil {
		return nil, err
	}

	return db, nil
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

		err = rows.Scan(&u.id, &u.cId, &u.itemId)

		if err != nil {
			return nil, err
		}

		reservations = append(reservations, u)
	}

	return reservations, nil
}

func setReservation(db *sql.DB, cId int, itemId int) (bool, error) {
	err := db.QueryRow(`INSERT INTO reservations(cId, itemId)
	VALUES($1, $2) RETURNING id`, cId, itemId).Scan()

	if err != nil {
		return false, err
	}else {
		return true, nil
	}
}

