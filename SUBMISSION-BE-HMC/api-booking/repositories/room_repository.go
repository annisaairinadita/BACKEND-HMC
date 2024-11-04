package repositories

import (
	"api-booking/models"
	"errors"

	"gorm.io/gorm"
)

type RoomRepository interface {
	Create(room models.Room) error
	GetAll(limit, offset int, roomType string, minPrice, maxPrice float64) ([]models.Room, error)
	CountRooms(roomType string, minPrice, maxPrice float64) (int64, error)
	GetRoomById(id int) (models.Room, error)
	GetRoomByRoomNumber(roomNumber string) (models.Room, error)
	UpdateRoom(room models.Room) error
	Delete(id int) error
}

type roomRepo struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepo{
		db: db,
	}
}

func (r *roomRepo) Create(room models.Room) error {
	err := r.db.Create(&room).Error
	if  err != nil {
		return err
	}
	return nil
}

func (r *roomRepo) GetAll(limit, offset int, roomType string, minPrice, maxPrice float64) ([]models.Room, error){
	var rooms []models.Room
	err := r.db.Where("room_type LIKE ? AND price_per_night BETWEEN ? AND ?", "%"+roomType+"%", minPrice, maxPrice).
        Limit(limit).Offset(offset).Find(&rooms).Error
	return rooms, err
}

func (r *roomRepo) CountRooms(roomType string, minPrice, maxPrice float64) (int64, error) {
    var count int64
    err := r.db.Model(&models.Room{}).Where("room_type LIKE ? AND price_per_night BETWEEN ? AND ?", "%"+roomType+"%", minPrice, maxPrice).Count(&count).Error
    return count, err
}

func (r *roomRepo) GetRoomById(id int) (models.Room, error) {
	var room models.Room
	result := r.db.First(&room, id)
	return room, result.Error
}

func (r *roomRepo) GetRoomByRoomNumber(roomNumber string) (models.Room, error) {
	var room models.Room
	result := r.db.Where("room_number = ?", roomNumber).First(&room)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return room, nil
		}
		return room, result.Error
	}
	return room, nil
}

func (r *roomRepo) UpdateRoom(room models.Room) error {
	err := r.db.Model(&models.Room{}).Where("id = ?", room.ID).Updates(&room).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roomRepo) Delete(id int) error {
	result := r.db.Delete(&models.Room{}, id).Error
	if result != nil {
		return result
	}
	return nil
}