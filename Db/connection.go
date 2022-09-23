package db

import (
	"database/sql"
	"fmt"
)

type DB struct {
	Db *sql.DB
}

func Init() *DB {
	return &DB{}
}

func (db *DB) CreateConnection() *DB {
	// Opening a driver typically will not attempt to connect to the database.
	database, err := sql.Open("mysql", "root:@/db_auth_golang")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		fmt.Println(err)
		return db
	}
	db.Db = database
	return db
}

func (db *DB) GetConnection() *DB {
	return db
}
