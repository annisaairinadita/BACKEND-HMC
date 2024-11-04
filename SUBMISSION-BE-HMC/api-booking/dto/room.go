package dto

import "api-booking/models"

type RoomRequest struct {
	RoomNumber    string  `json:"room_number"`
	RoomType      string  `json:"room_type"`
	PricePerNight float64 `json:"price_per_night"`
}

type RoomResponse struct {
	ID            int     `json:"id"`
	RoomNumber    string  `json:"room_number"`
	RoomType      string  `json:"room_type"`
	PricePerNight float64 `json:"price_per_night"`
}

func NewRoomResponse(room models.Room) RoomResponse {
	return RoomResponse{
		ID:            room.ID,
		RoomNumber:    room.RoomNumber,
		RoomType:      room.RoomType,
		PricePerNight: room.PricePerNight,
	}
}

func NewListRooms(rooms []models.Room) []RoomResponse {
	listRooms := []RoomResponse{}
	for _, room := range rooms {
		roomDetail := NewRoomResponse(room)
		listRooms = append(listRooms, roomDetail)
	}
	return listRooms
}