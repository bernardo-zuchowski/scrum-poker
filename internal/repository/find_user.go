package repository

import (
	"multipoker/internal/models"
)

func FindUser(userId int, roomId int) *models.User {
	room := FindRoom(roomId)

	for _, user := range room.Users {
		if user.ID == userId {
			return user
		}
	}

	return nil
}
