package model

import "time"

// Building struct
type Building struct {
	Id        int16     `json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Area      string    `json:"area"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Buildings struct
type Buildings struct {
	Buildings []Building `json:"buildings"`
}
