package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// Building struct
type Building struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	LocationId   int64     `json:"location_id"`
	LocationName string    `json:"location_name"`
	Address      string    `json:"address"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Buildings struct
type Buildings struct {
	Buildings []*Building
}

func GetAllBuildings() []*Building {
	// query all data
	rows, e := migration.DbPool.Query(context.Background(), "select * from buildings")
	if e != nil {
		return nil
	}
	defer rows.Close()

	// declare empty post variable
	buildings := make([]*Building, 0)
	// iterate over rows
	for rows.Next() {
		building := new(Building)
		e = rows.Scan(&building.Id, &building.Name, &building.LocationId, &building.Address, &building.CreatedAt, &building.UpdatedAt)
		location := migration.DbPool.QueryRow(context.Background(), "select name from locations where id = $1", &building.LocationId)
		err := location.Scan(&building.LocationName)
		if err != nil {
			return nil
		}
		if e != nil {
			fmt.Println("Failed to get buildings record :", e)
			return []*Building{}
		}
		buildings = append(buildings, building)
	}
	return buildings
}

func (building *Building) CreateBuilding() error {
	dt := time.Now()
	query := "INSERT INTO buildings (name, location_id, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(context.Background(), query, &building.Name, &building.LocationId, &building.Address, dt, dt)
	err := d.Scan(&building.Id, &building.CreatedAt, &building.UpdatedAt)
	if err != nil {
		return err
	}
	location := migration.DbPool.QueryRow(context.Background(), "select name from locations where id = $1", &building.LocationId)
	err = location.Scan(&building.LocationName)
	if err != nil {
		return err
	}
	return nil
}

func GetBuildingByID(buildingId int) Building {
	building := Building{}
	rows := migration.DbPool.QueryRow(context.Background(), "select buildings.id, buildings.name, buildings.location_id, locations.name as location_name, buildings.address, buildings.created_at, buildings.updated_at from buildings LEFT JOIN locations on buildings.location_id = locations.id where buildings.id = $1", buildingId)
	err := rows.Scan(&building.Id, &building.Name, &building.LocationId, &building.LocationName, &building.Address, &building.CreatedAt, &building.UpdatedAt)
	if err != nil {
		fmt.Println("Failed to get locations record :", err)
		return Building{}
	}
	return building
}
