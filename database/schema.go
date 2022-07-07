package database

import "fmt"

// CreateProductTable ...
func CreateRoleTable() {

	r, err := DB.Query(`CREATE TABLE IF NOT EXISTS roles (
        id serial PRIMARY KEY,
        name VARCHAR ( 50 ) UNIQUE NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.Columns())
	}

}
