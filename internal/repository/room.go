package repository

import "multipoker/internal/models"

var Rooms []*models.Room
var roomID int

func GetNextRoomID() int {
	roomID++
	return roomID
}
