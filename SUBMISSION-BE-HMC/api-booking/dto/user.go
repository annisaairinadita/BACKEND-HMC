package dto

import "api-booking/models"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetail struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

func NewUserDetail(user models.User) UserDetail {
	return UserDetail{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
	}
}

func NewListUsers(users []models.User) []UserDetail {
	listUser := []UserDetail{}
	for _, user := range users {
		userDetail := NewUserDetail(user)
		listUser = append(listUser, userDetail)
	}
	return listUser
}
