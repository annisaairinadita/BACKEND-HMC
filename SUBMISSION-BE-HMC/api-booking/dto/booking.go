package dto

import (
	"api-booking/models"
	"fmt"
	"time"
)

type BookingRequest struct {
	RoomID   int    `json:"room_id" validate:"required"`
	CheckIn  string `json:"check_in" validate:"required"`
	CheckOut string `json:"check_out" validate:"required"`
}

type BookingResponse struct {
	ID           int       `json:"id"`
	UserID       uint      `json:"user_id"`
	RoomID       int       `json:"room_id"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	TotalPrice   float64   `json:"total_price"`
}

type BookingListResponse struct {
	Bookings []BookingResponse `json:"bookings"`
}

func (b *BookingRequest) ToBooking() (*models.Booking, error) {
	checkIn, err := time.Parse("2006-01-02T15:04:05Z07:00", b.CheckIn)
	if err != nil {
		return nil, fmt.Errorf("invalid check-in time format")
	}
	checkOut, err := time.Parse("2006-01-02T15:04:05Z07:00", b.CheckOut)
	if err != nil {
		return nil, fmt.Errorf("invalid check-out time format")
	}
	return &models.Booking{
		RoomID:       b.RoomID,
		CheckInDate:  checkIn,
		CheckOutDate: checkOut,
	}, nil
}

func (b *BookingResponse) FromBooking(booking *models.Booking) {
	b.ID = booking.ID
	b.UserID = booking.UserID
	b.RoomID = booking.RoomID
	b.CheckInDate = booking.CheckInDate
	b.CheckOutDate = booking.CheckOutDate
	b.TotalPrice = booking.TotalPrice
}

func (b *BookingListResponse) FromBookings(bookings []models.Booking)  {
	b.Bookings = make([]BookingResponse, len(bookings))
	for i, booking := range bookings {
		b.Bookings[i].FromBooking(&booking)
	}
}