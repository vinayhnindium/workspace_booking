package database

// import (
// 	"database/sql"
// 	"fmt"
// )

// // Database instance
// var DB *sql.DB

// // Connect function
// func Connect() error {
// 	var err error
// 	DB, err = sql.Open("postgres", "host=localhost port=5432 user=mahesh dbname=workspace_booking sslmode=disable")

// 	if err != nil {
// 		return err
// 	}

// 	err = DB.Ping()

// 	if err != nil {
// 		return err
// 	}

// 	CreateRoleTable()

// 	fmt.Println("Connection Opened to Database")
// 	return nil
// }

import (
	"context"
	"workspace_booking/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

var dbPool *pgxpool.Pool

func GetDbConnectionPool() *pgxpool.Pool {
	if dbPool != nil {
		return dbPool
	}

	psqlconn := config.GetDBConnectionURL()
	println(psqlconn)
	db, err := pgxpool.Connect(context.Background(), psqlconn)

	// open database
	checkError(err)
	dbPool = db

	CreateRoleTable()

	return dbPool
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
