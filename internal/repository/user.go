package repository

import "multipoker/internal/models"

var userID int
var Users []*models.User

func GetNextUserID() int {
	userID++
	return userID
}
