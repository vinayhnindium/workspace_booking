package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// Role struct
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

// GetAllRoles will fetch all the roles from roles table
func GetAllRoles() []Role {
	fmt.Println("Fetching")

	rows, err := migration.DbPool.Query(context.Background(), "SELECT * FROM roles")
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

// InsertRole will create the role record in db
func (r *Role) InsertRole() error {

	dt := time.Now()
	query := "INSERT INTO roles (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(context.Background(), query, r.Name, dt, dt)
	err := d.Scan(&r.Id, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		return err
	}

	return nil

}
