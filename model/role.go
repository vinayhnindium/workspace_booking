package model

import "time"

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
