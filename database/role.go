package database

import (
	"context"
	"fmt"
	"time"
)

// role struct
type Role struct {
	Name      string    `json:"name"`
	Id        int16     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Roles struct
type Roles struct {
	Roles []Role `json:"roles"`
}

func GetAllRoles() []Role {
	fmt.Println("Fetching")

	rows, err := dbPool.Query(context.Background(), "SELECT * FROM roles")
	if err != nil {
		return nil
	}

	defer rows.Close()

	roles := make([]Role, 0)
	for rows.Next() {
		r := Role{}
		err = rows.Scan(&r.Id, &r.Name, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			fmt.Println("Failed", err)
			return []Role{}
		}

		roles = append(roles, Role{Id: r.Id, Name: r.Name, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt})
	}

	return roles
}

func CreateRole(name string) (Role, error) {

	row, err := dbPool.Query(context.Background(), "INSERT INTO roles (name) VALUES ($1)", name)
	if err != nil {
		fmt.Println("Failed", err)
		return Role{}, err
	}
	fmt.Println("************", &row)
	r := Role{}
	row.Scan(&r.Id, &r.Name)

	return r, nil

}
