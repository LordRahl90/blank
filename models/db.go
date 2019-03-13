package models

import "database/sql"

//Database struct.
type Database struct {
	DB *sql.DB
}

//Close - function to close the database connection
func (db Database) Close() {
	db.DB.Close()
}
