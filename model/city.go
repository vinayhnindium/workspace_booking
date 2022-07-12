package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// Building struct
type City struct {
	Id        int16     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Buildings struct
type Cities struct {
	Cities []City `json:"cities"`
}

func GetAllCities() []City {
	rows, e := migration.DbPool.Query(context.Background(), "select * from cities")
	defer rows.Close()

	// declare empty post variable
	cities := make([]City, 0)
	// iterate over rows
	for rows.Next() {
		city := City{}
		e = rows.Scan(&city.Id, &city.Name, &city.CreatedAt, &city.UpdatedAt)
		if e != nil {
			fmt.Println("Failed to get buildings record :", e)
			return []City{}
		}
		cities = append(cities, City{Id: city.Id, Name: city.Name, CreatedAt: city.CreatedAt, UpdatedAt: city.UpdatedAt})
	}
	return cities
}

func (city *City) CreateCity() error {
	dt := time.Now()
	query := "INSERT INTO cities (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(context.Background(), query, &city.Name, dt, dt)
	d.Scan(&city.Id, &city.CreatedAt, &city.UpdatedAt)
	return nil
}
