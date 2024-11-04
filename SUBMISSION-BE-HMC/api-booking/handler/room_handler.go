package handler

import (
	"api-booking/models"
	"api-booking/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomHandler interface {
	Create(c *gin.Context)
	GetAllRoom(c *gin.Context)
	GetRoomById(c *gin.Context)
	GetRoomByRoomNumber(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type roomHandler struct {
	serv services.RoomService
}

func NewRoomHandler(serv services.RoomService) RoomHandler {
	return &roomHandler{serv: serv}
}

func (r *roomHandler) Create(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := r.serv.Create(room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success create room"})
}

func (r *roomHandler) GetAllRoom(c *gin.Context) {
	roomType := c.Query("type")
    minPrice, _ := strconv.ParseFloat(c.Query("min_price"), 64)
    maxPrice, _ := strconv.ParseFloat(c.Query("max_price"), 64)
    limit, _ := strconv.Atoi(c.Query("limit"))
    offset, _ := strconv.Atoi(c.Query("offset"))
	rooms, err := r.serv.GetAll(limit, offset, roomType, minPrice, maxPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalRooms, _ := r.serv.CountRooms(roomType, minPrice, maxPrice)
	
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success get all room",
		"data":    rooms,
		"total": totalRooms,
	})
}

func (r *roomHandler) GetRoomById(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	room, err := r.serv.GetRoomById(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success get room by id",
		"data":    room,
	})
}

func (r *roomHandler) GetRoomByRoomNumber(c *gin.Context) {
	roomNumber := c.Param("roomNumber")
	room, err := r.serv.GetRoomByRoomNumber(roomNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success get room by room number",
		"data":    room,
	})
}

func (r *roomHandler) Update(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	room.ID = num

	err = r.serv.UpdateRoom(room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})	
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success update room"})
}

func (r *roomHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return
	}

	err = r.serv.Delete(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success delete room"})
}
