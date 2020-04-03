package pgstore

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // register its drivers with the database/sql package
)

func connString() string {
	return os.ExpandEnv("host=$DB_HOST port=5432 user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable")
}

// OpenDB returns open database, configured from env
func OpenDB() (db *sql.DB, err error) {

	// TODO: Set max conections: check the optimal setting

	db, err = sql.Open("postgres", connString())
	return
}

// MustOpenDB returns open database, configured from env. It check that db is alive and panics on any error
func MustOpenDB() *sql.DB {
	db, err := OpenDB()
	if err != nil {
		panic(fmt.Errorf("OpenDB: %w", err))
	}
	if err = db.Ping(); err != nil {
		panic(fmt.Errorf("OpenDB.Ping: %w", err))
	}
	return db
}
