package handler

import (
	"api-booking/dto"
	"api-booking/helpers"
	"api-booking/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingHandler interface {
	GetAllBooking(c *gin.Context)
	GetBookingById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type bookingHandler struct {
	serv services.BookingService
}

func NewBookingHandler(serv services.BookingService) BookingHandler {
	return &bookingHandler{serv: serv}
}

func (b *bookingHandler) GetAllBooking(c *gin.Context) {
	bookings, err := b.serv.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bo := dto.BookingListResponse{}
	bo.FromBookings(bookings)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success get all booking",
		"data":    bo,
	})
}

func (b *bookingHandler) GetBookingById(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}
	booking, err := b.serv.GetByID(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	bb := dto.BookingResponse{}
	bb.FromBooking(&booking)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success get booking by id",
		"data":    booking,
	})
}

func (b *bookingHandler) Create(c *gin.Context) {
	var booking dto.BookingRequest
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	id, _, err := helpers.GetUserClaims(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bo, err := booking.ToBooking()
	bo.UserID = uint(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = b.serv.Create(*bo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "success create booking",
	})
}

func (b *bookingHandler) Update(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}
	var booking dto.BookingRequest
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	bo, err := booking.ToBooking() // Handle error properly
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bo.ID = num
	err = b.serv.Update(bo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(*bo)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success update booking",
	})
}

func (b *bookingHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}
	err = b.serv.Delete(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success delete booking",
	})
}
