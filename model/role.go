package model

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

	rows, err := DbPool.Query(context.Background(), "SELECT * FROM roles")
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

func (r *Role) CreateRole(name string) error {

	dt := time.Now()
	query := "INSERT INTO roles (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	d := DbPool.QueryRow(context.Background(), query, name, dt, dt)
	d.Scan(&r.Id, &r.CreatedAt, &r.UpdatedAt)

	return nil

}

// db.QueryRow("INSERT INTO users (name, email, encrypted_password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at",
// u.Name, u.Email, u.Password, dt, dt).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
