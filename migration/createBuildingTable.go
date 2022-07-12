package migration

import (
	"context"
	"fmt"
)

func CreateBuildingTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS buildings (
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
		fmt.Println(r)
	}

}
