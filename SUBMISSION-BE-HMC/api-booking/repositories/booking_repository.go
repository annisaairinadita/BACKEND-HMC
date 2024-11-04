package repositories

import (
	"api-booking/models"
	"fmt"

	"gorm.io/gorm"
)

type BookingRepository interface {
	GetAll() ([]models.Booking, error)
	GetById(id int) (models.Booking, error)
	Create(booking models.Booking) error
	Update(booking *models.Booking) error
	Delete(id int) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{
		db: db,
	}
}

func (r *bookingRepository) GetAll() ([]models.Booking, error) {
	var bookings []models.Booking
	return bookings, r.db.Find(&bookings).Error
}

func (r *bookingRepository) GetById(id int) (models.Booking, error) {
	var booking models.Booking
	return booking, r.db.First(&booking, id).Error
}

func (r *bookingRepository) Create(booking models.Booking) error {
	return r.db.Create(&booking).Error
}

func (r *bookingRepository) Update(booking *models.Booking) error {
	fmt.Println("iniii repooooo ",*booking)
	err := r.db.Model(&models.Booking{}).Where("id = ?", booking.ID).Updates(&booking).Error
	if err != nil{
		return err
	}
	return nil
}

func (r *bookingRepository) Delete(id int) error {
	return r.db.Delete(&models.Booking{}, id).Error
}