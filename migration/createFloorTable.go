package migration

import (
	"context"
	"fmt"
)

// CreateRoleTable ...
func CreateFloorTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS floors (
		id serial PRIMARY KEY,
		name VARCHAR ( 50 ) UNIQUE NOT NULL,
		total_workspace INTEGER NOT NULL,
		building_id INTEGER REFERENCES buildings (id),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
	`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
