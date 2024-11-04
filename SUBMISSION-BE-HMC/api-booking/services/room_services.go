package services

import (
	"api-booking/models"
	"api-booking/repositories"

	"github.com/go-playground/validator/v10"
)

type RoomService interface {
	Create(room models.Room) error
	GetAll(limit, offset int, roomType string, minPrice, maxPrice float64) ([]models.Room, error)
	CountRooms(roomType string, minPrice, maxPrice float64) (int64, error) 
	GetRoomById(id int) (models.Room, error)
	GetRoomByRoomNumber(roomNumber string) (models.Room, error)
	UpdateRoom(room models.Room) error
	Delete(id int) error
}

type roomService struct {
	roomRepo repositories.RoomRepository
}

func NewRoomService(roomRepo repositories.RoomRepository) RoomService {
	return &roomService{
		roomRepo: roomRepo,
	}
}

func (s *roomService) Create(room models.Room) error {
	validator := validator.New()
	err := validator.Struct(room)
	if err != nil {
		return err
	}
	return s.roomRepo.Create(room)
}

func (s *roomService) GetAll(limit, offset int, roomType string, minPrice, maxPrice float64) ([]models.Room, error) {
	rooms, err := s.roomRepo.GetAll(limit, offset, roomType, minPrice, maxPrice)
	if err != nil {
		return []models.Room{}, err
	}
	return rooms, nil
}

func (s *roomService) CountRooms(roomType string, minPrice, maxPrice float64) (int64, error) {
    return s.roomRepo.CountRooms(roomType, minPrice, maxPrice)
}

func (s *roomService) GetRoomById(id int) (models.Room, error) {
	room, err := s.roomRepo.GetRoomById(id)
	if err != nil {
		return models.Room{}, err
	}
	return room, nil
}

func (s *roomService) GetRoomByRoomNumber(roomNumber string) (models.Room, error) {
	room, err := s.roomRepo.GetRoomByRoomNumber(roomNumber)
	if err != nil {
		return models.Room{}, err
	}
	return room, nil
}

func (s *roomService) UpdateRoom(room models.Room) error {
	validator := validator.New()
	err := validator.Struct(room)
	if err != nil {
		return err
	}
	return s.roomRepo.UpdateRoom(room)
}

func (s *roomService) Delete(id int) error {
	return s.roomRepo.Delete(id)
}