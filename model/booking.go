package model

import (
	"context"
	"time"
	"workspace_booking/migration"
)

type Booking struct {
	Id               int16     `json:"id"`
	CityId           int16     `json:"city_id"`
	LocationId       int16     `json:"location_id"`
	BuildingId       int16     `json:"building_id"`
	FloorId          int16     `json:"floor_id"`
	FromDate         string    `json:"from_date"`
	ToDate           string    `json:"to_date"`
	Purpose          string    `json:"purpose"`
	UserId           int16     `json:"user_id"`
	WorkspacesBooked int16     `json:"workspaces_booked"`
	UserIds          []int16   `json:"user_ids"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// InsertBooking will create the booking record in db
func (b *Booking) InsertBooking() error {
	dt := time.Now()
	query := "INSERT INTO bookings (city_id, location_id, building_id, floor_id, from_date, to_date, purpose, " +
		"user_id, workspaces_booked, created_at, updated_at)" +
		" VALUES ($1, $2, $3, $4,$5, $6, $7, $8, $9, $10, $11) RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(
		context.Background(), query, b.CityId, b.LocationId, b.BuildingId, b.FloorId, b.FromDate, b.ToDate,
		b.Purpose, b.UserId, b.WorkspacesBooked, dt, dt,
	)
	err := d.Scan(&b.Id, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return err
	}

	return nil

}
