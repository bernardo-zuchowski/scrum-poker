package repository

import (
	"multipoker/internal/models"
)

func GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	for _, room := range Rooms {
		users = append(users, room.Users...)
	}

	return users, nil
}
