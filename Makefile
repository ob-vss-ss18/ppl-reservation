DATABASE_NAME = heroku-go-db-example

run: .env create bin/ppl-reservation
	@PATH="$(PWD)/bin:$(PATH)" heroku local

.env:
	cp .env.dev .env

bin/ppl-reservation: main.go
	go build -o bin/ppl-reservation main.go

create:
	@psql -lqt | cut -d \| -f 1 | grep -qw $(DATABASE_NAME) || createdb $(DATABASE_NAME)

drop:
	dropdb $(DATABASE_NAME)

clean:
	rm -rf bin

