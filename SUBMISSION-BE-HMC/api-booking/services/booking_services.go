package services

import (
	"api-booking/models"
	"api-booking/repositories"
	"fmt"
	"time"
)

type BookingService interface {
	Create(booking models.Booking) error
	GetAll() ([]models.Booking, error)
	GetByID(id int) (models.Booking, error)
	Update(booking *models.Booking) error
	Delete(id int) error
}

type bookingService struct {
	bookingRepo repositories.BookingRepository
	roomRepo    repositories.RoomRepository
}

func NewBookingService(bookingRepo repositories.BookingRepository, roomRepo repositories.RoomRepository) BookingService {
	return &bookingService{
		bookingRepo: bookingRepo,
		roomRepo:    roomRepo,
	}
}

func (s *bookingService) Create(booking models.Booking) error {
	room, err := s.roomRepo.GetRoomById(booking.RoomID)
	if err != nil {
		return err
	}

	booking.TotalPrice = room.PricePerNight * booking.CheckOutDate.Sub(booking.CheckInDate).Hours() / 24

	now := time.Now()
	var penalty float64

	if now.After(booking.CheckOutDate.Add(24 * time.Hour)) {
	
		overdueDuration := now.Sub(booking.CheckOutDate)
		overdueHours := int(overdueDuration.Hours())

		penalty = float64(overdueHours) * 50.0 // Denda 50 per jam keterlambatan
	}

	booking.TotalPrice += penalty

	err = s.bookingRepo.Create(booking)
	return err
}

func (s *bookingService) GetAll() ([]models.Booking, error) {
	return s.bookingRepo.GetAll()
}

func (s *bookingService) GetByID(id int) (models.Booking, error) {
	return s.bookingRepo.GetById(id)
}

func (s *bookingService) Update(booking *models.Booking) error {
	existingBooking, err := s.bookingRepo.GetById(booking.ID)
	if err != nil {
		return err 
	}
	existingBooking.RoomID = booking.RoomID

	existingBooking.CheckInDate = booking.CheckInDate
	existingBooking.CheckOutDate = booking.CheckOutDate

	room, err := s.roomRepo.GetRoomById(existingBooking.RoomID)
	if err != nil {
		return err
	}
	existingBooking.TotalPrice = room.PricePerNight * existingBooking.CheckOutDate.Sub(existingBooking.CheckInDate).Hours() / 24

	now := time.Now()
	var penalty float64

	if now.After(existingBooking.CheckOutDate.Add(24 * time.Hour)) {
	
		overdueDuration := now.Sub(existingBooking.CheckOutDate)
		overdueHours := int(overdueDuration.Hours())

		penalty = float64(overdueHours) * 50.0 
	}

	existingBooking.TotalPrice += penalty
	fmt.Println("ini serviceeeeee ", existingBooking)

	err = s.bookingRepo.Update(&existingBooking)
	return err
}
func (s *bookingService) Delete(id int) error {
	return s.bookingRepo.Delete(id)
}
