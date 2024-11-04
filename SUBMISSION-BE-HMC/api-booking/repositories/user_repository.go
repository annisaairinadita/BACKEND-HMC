package repositories

import (
	"api-booking/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) error
	GetAll() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	UpdateUser(user models.User) error
	Delete(id uint) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(user models.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) Delete(id uint) error {
	err := u.db.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) GetAll() ([]models.User, error) {
	var users []models.User
	result := u.db.Find(&users)
	return users, result.Error
}

func (u *userRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := u.db.Debug().Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, result.Error
	}
	return user, nil
}

func (u *userRepo) GetUserById(id uint) (models.User, error) {
	var user models.User
	result := u.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, nil
		}
		return user, result.Error
	}
	return user, nil
}

func (u *userRepo) UpdateUser(user models.User) error {
	result := u.db.Save(&user)
	return result.Error
}
