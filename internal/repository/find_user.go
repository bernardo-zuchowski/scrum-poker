package repository

import (
	"multipoker/internal/models"
)

func FindUser(id int) *models.User {
	for _, user := range Users {
		if user.ID == id {
			return user
		}
	}

	return nil
}
