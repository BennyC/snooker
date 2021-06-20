package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDatabase(c string) *sql.DB {
	db, err := sql.Open("postgres", c)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
