package repository

import "multipoker/internal/models"

var Users []*models.User
var userID int

func GetNextUserID() int {
	userID++
	return userID
}
