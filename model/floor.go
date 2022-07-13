package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// Building struct
type Floor struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	TotalWorkSpace int       `json:"total_workspace"`
	BuildingId     int64     `json:"building_id"`
	BuildingName   string    `json:"building_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Floors struct
type Floors struct {
	Floors []*Floor
}

func GetAllFloors() []*Floor {
	rows, _ := migration.DbPool.Query(context.Background(), "select floors.id, floors.name, floors.total_workspace, floors.building_id, buildings.name as building_name, floors.created_at, floors.updated_at from floors LEFT JOIN buildings ON floors.building_id = buildings.id")
	defer rows.Close()

	// declare empty post variable
	floors := make([]*Floor, 0)
	// iterate over rows
	for rows.Next() {
		floor := new(Floor)
		err := rows.Scan(&floor.Id, &floor.Name, &floor.TotalWorkSpace, &floor.BuildingId, &floor.BuildingName, &floor.CreatedAt, &floor.UpdatedAt)
		if err != nil {
			fmt.Println("Failed to get floors record :", err)
			return []*Floor{}
		}
		floors = append(floors, floor)
	}
	return floors
}

func (floor *Floor) CreateFloor() error {
	dt := time.Now()
	query := "INSERT INTO floors (name, total_workspace, building_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at"
	location := migration.DbPool.QueryRow(context.Background(), "select name from buildings where id = $1", &floor.BuildingId)
	location.Scan(&floor.BuildingName)
	d := migration.DbPool.QueryRow(context.Background(), query, &floor.Name, &floor.TotalWorkSpace, &floor.BuildingId, dt, dt)
	d.Scan(&floor.Id, &floor.CreatedAt, &floor.UpdatedAt)
	return nil
}

func GetFloorByID(floorId int) Floor {
	floor := Floor{}
	// rows := migration.DbPool.QueryRow(context.Background(), "select floors.id, floors.name, floors.total_workspace, floors.building_id, buildings.name as building_name, floors.created_at, floors.updated_at from floors LEFT JOIN buildings ON floors.building_id = buildings.id where floors.id = $1", floorId)
	rows := migration.DbPool.QueryRow(context.Background(), "select * from floors where id = $1", floorId)
	err := rows.Scan(&floor.Id, &floor.Name, &floor.TotalWorkSpace, &floor.BuildingId, &floor.CreatedAt, &floor.UpdatedAt)
	if err != nil {
		fmt.Println("Failed to get floors record :", err)
		return Floor{}
	}
	return floor
}
