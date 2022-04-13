package repository

import (
	"multipoker/internal/models"
)

func GetAllRooms() ([]*models.Room, error) {
	return Rooms, nil
}
