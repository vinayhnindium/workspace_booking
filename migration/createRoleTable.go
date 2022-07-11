package migration

import (
	"context"
	"fmt"
)

// CreateRoleTable ...
func CreateRoleTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS roles (
        id serial PRIMARY KEY,
        name VARCHAR ( 50 ) UNIQUE NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func CreateBuildingTable() {

	r, err := DB.Query(`CREATE TABLE buildings (
        id serial PRIMARY KEY,
        name VARCHAR ( 50 ) NOT NULL,
        city VARCHAR ( 50 ) NOT NULL,
        area VARCHAR ( 50 ) NOT NULL,
        address VARCHAR ( 50 ) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.Columns())
	}

}
