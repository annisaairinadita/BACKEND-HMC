package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	PhoneNumber string         `json:"PhoneNumber"`
	Role        string         `json:"role" enum:"admin,user"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Room struct {
    ID            int            `json:"id"`
    RoomNumber    string         `json:"room_number" validate:"required"` 
    RoomType      string         `json:"room_type" validate:"required"`   
    PricePerNight float64        `json:"price_per_night" validate:"required,gte=0"` 
    CreatedAt     time.Time      `json:"created_at"`
    UpdatedAt     time.Time      `json:"updated_at"`
    DeletedAt     gorm.DeletedAt `gorm:"index" yjson:"deleted_at"`
}

type Booking struct {
	ID           int            `json:"id"`
	UserID       uint           `json:"user_id"` 
	RoomID       int            `json:"room_id"` 
	CheckInDate  time.Time      `json:"check_in_date"`
	CheckOutDate time.Time      `json:"check_out_date"`
	TotalPrice   float64        `json:"total_price"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}