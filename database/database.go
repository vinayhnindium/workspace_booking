package database

import (
	"database/sql"
	"fmt"
)

// Database instance
var DB *sql.DB

// Connect function
func Connect() error {
	var err error
	DB, err = sql.Open("postgres", "host=localhost port=5432 user=mahesh dbname=workspace_booking sslmode=disable")
	fmt.Println("DB******", DB)
	if err != nil {
		return err
	}

	err = DB.Ping()

	if err != nil {
		return err
	}

	CreateRoleTable()
	// CreateTables()
	fmt.Println("Connection Opened to Database")
	return nil
}
