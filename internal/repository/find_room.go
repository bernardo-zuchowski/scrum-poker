package repository

import (
	"multipoker/internal/models"
)

func FindRoom(id int) *models.Room {
	for _, room := range Rooms {
		if room.ID == id {
			return room
		}
	}

	return nil
}
