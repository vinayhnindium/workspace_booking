package model

import (
	"context"
	"time"
	"workspace_booking/migration"
)

// BookingParticipant struct
type BookingParticipant struct {
	Id        int16     `json:"id"`
	BookingId int16     `json:"booking_id"`
	UserId    int16     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (bp *BookingParticipant) CreateBookingParticipant() error {
	dt := time.Now()
	query := "INSERT INTO booking_participants (booking_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4) " +
		"RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(context.Background(), query, &bp.BookingId, &bp.UserId, dt, dt)
	err := d.Scan(&bp.Id, &bp.CreatedAt, &bp.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func BulkInsertBookingParticipant(bookingId int16, userIds []int16) error {
	for _, userId := range userIds {
		bookingParticipant := new(BookingParticipant)
		bookingParticipant.BookingId = bookingId
		bookingParticipant.UserId = userId
		err := bookingParticipant.CreateBookingParticipant()
		if err != nil {
			return err
		}
	}
	return nil
}
