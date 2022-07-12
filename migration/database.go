package migration

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"workspace_booking/config"
)

var DbPool *pgxpool.Pool

func GetDbConnectionPool() *pgxpool.Pool {
	if DbPool != nil {
		return DbPool
	}

	psqlconn := config.GetDBConnectionURL()
	println(psqlconn)
	db, err := pgxpool.Connect(context.Background(), psqlconn)

	// open database
	checkError(err)
	DbPool = db

	// Dont change the order here
	CreateRoleTable()
	CreateBookingsTable()

	return DbPool
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
