package main

import (
	"database/sql"
	_ "github.com/lib/pq"
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
	  date_from DATE,
	  date_to DATE
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
	err := db.QueryRow(`INSERT INTO reservations(cId, itemId, to_date(date_from, "DD.MM.YYYY"), to_date(date_to, "DD.MM.YYYY"))
	VALUES($1, $2, $3, $4) RETURNING id`, cId, itemId, date_from, date_to).Scan(&id)

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
