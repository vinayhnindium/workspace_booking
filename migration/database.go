package migration

import (
	"context"
	"workspace_booking/config"

	"github.com/jackc/pgx/v4/pgxpool"
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
	CreateBuildingTable()
	CreateCityTable()
	CreateLocationTable()
	CreateFloorTable()

	return DbPool
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
