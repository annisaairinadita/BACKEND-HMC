package services 
import (
	"api-booking/helpers"
	"api-booking/models"
	"api-booking/repositories"
	"errors"
)

type UserService interface {
	Login(user models.User) (string, error)
	GetAll() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	Register(user models.User) (models.User, error)
	Update(user models.User) error 
	Delete(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(rp repositories.UserRepository) UserService {
	return &userService{
		repo: rp,
	}
}

func (u *userService) Delete(id uint) error {
	err := u.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) GetAll() ([]models.User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userService) GetUserById(id uint) (models.User, error) {
	user, err := u.repo.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userService) Login(user models.User) (string, error) {
	existingUser, err := u.repo.GetUserByEmail(user.Email)

	if err != nil {
		return "", err
	}

	ok, err := helpers.ComparePass([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("invalid password")

	}
	if existingUser.ID == 0 {
		return "", errors.New("user not found")
	}
	token, err := helpers.CreateTokenJWT(int(existingUser.ID), existingUser.Role)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *userService) Register(user models.User) (models.User, error) {
	existingUser, err := u.repo.GetUserByEmail(user.Email)
	if err == nil && existingUser.ID != 0 {
		return models.User{}, errors.New("email already registered")
	}

	user.Role = "user"

	hassPass, err := helpers.HassPass(user.Password)
	if err != nil {
		return models.User{}, err //
	}
	user.Password = hassPass
	if err := u.repo.Create(user); err != nil {
		return models.User{}, err //
	}

	return user, nil
}

func (u *userService) Update(user models.User) error {
	existingUser, err := u.repo.GetUserById(user.ID)
	if err != nil {
		return err
	}
	if existingUser.ID == 0 {
		return errors.New("user not found")
	}

	if user.Password != "" && user.Password != existingUser.Password {
		hashedPassword, err := helpers.HassPass(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}

	if err := u.repo.UpdateUser(user); err != nil {
		return err
	}

	return nil
}