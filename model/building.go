package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

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

func GetAllBuildings() []Building {
	// query all data
	rows, e := migration.DbPool.Query(context.Background(), "select * from buildings")
	if e != nil {
		return nil
	}
	defer rows.Close()

	// declare empty post variable
	buildings := make([]Building, 0)
	// iterate over rows
	for rows.Next() {
		building := Building{}
		e = rows.Scan(&building.Id, &building.Name, &building.City, &building.Area, &building.Address, &building.CreatedAt, &building.UpdatedAt)
		if e != nil {
			fmt.Println("Failed to get buildings record :", e)
			return []Building{}
		}
		buildings = append(buildings, Building{Id: building.Id, Name: building.Name, City: building.City, Area: building.Area, Address: building.Address, CreatedAt: building.CreatedAt, UpdatedAt: building.UpdatedAt})
	}
	return buildings
}

func (building *Building) CreateBuilding() error {
	dt := time.Now()
	query := "INSERT INTO buildings (name, city, area, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(context.Background(), query, &building.Name, &building.City, &building.Area, &building.Address, dt, dt)
	d.Scan(&building.Id, &building.CreatedAt, &building.UpdatedAt)
	return nil
}
