package testutils

import (
	"database/sql"

	"github.com/fudge/snooker/internal/storage/postgres"

	"github.com/golang-migrate/migrate/v4"
	mpsql "github.com/golang-migrate/migrate/v4/database/postgres"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const connString = "host=127.0.0.1 port=54322 user=postgres password=example dbname=snooker_test sslmode=disable"

func NewDatabase() *sql.DB {
	return postgres.New(connString)
}

func Migrate(db *sql.DB) func() {
	driver, err := mpsql.WithInstance(db, &mpsql.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://../../db/migrations", "postgres", driver)
	if err != nil {
		panic(err)
	}

	m.Up()
	return func() {
		m.Down()
	}
}
