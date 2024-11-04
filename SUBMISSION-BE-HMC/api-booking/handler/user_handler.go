package handler

import (
	"api-booking/dto"
	"api-booking/helpers"
	"api-booking/models"
	"api-booking/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProfile(c *gin.Context)
	GetAllUser(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Me(c *gin.Context)
}

type userHandler struct {
	serv services.UserService
}

func NewUserHandler(serv services.UserService) UserHandler {
	return &userHandler{serv: serv}
}

// Delete implements UserHandler.
func (u *userHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	err = u.serv.Delete(uint(num))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})

}

func (u *userHandler) GetAllUser(c *gin.Context) {
	users, err := u.serv.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	respUser := dto.NewListUsers(users)

	c.JSON(http.StatusOK, respUser)
}

func (u *userHandler) GetProfile(c *gin.Context) {
	id_user, role, err := helpers.GetUserClaims(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
    if role != "admin" && strconv.Itoa(id_user) != id {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to access this data"})
        return
    }
	userID := id
	num, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := u.serv.GetUserById(uint(num))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	userResp := dto.NewUserDetail(user)

	c.JSON(http.StatusOK, userResp)
}

func (u *userHandler) Login(c *gin.Context) {
	var user dto.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userModel := models.User{
		Email:    user.Email,
		Password: user.Password,
	}

	token, err := u.serv.Login(userModel)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

// Me implements UserHandler.
func (u *userHandler) Me(c *gin.Context) {
	panic("unimplemented")
}

// Register implements UserHandler.
func (u *userHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := u.serv.Register(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success create user",
	})
}

// Update implements UserHandler.
func (u *userHandler) Update(c *gin.Context) {
	id_user, role, err := helpers.GetUserClaims(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
    if role != "admin" && strconv.Itoa(id_user) != id {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to access this data"})
        return
    }

	userID, err := strconv.ParseUint(id, 10, 32) // Konversi id ke uint64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user.ID = uint(userID) // Set ID pengguna yang ingin diperbarui

	err = u.serv.Update(user) // Panggil fungsi update dari service
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update user",
	})
}