package repository

import (
	"multipoker/internal/dto"
	"multipoker/internal/models"
)

func CreateUser(data *dto.CreateUserDTO) (*models.User, error) {
	room := FindRoom(data.RoomID)

	user := &models.User{
		ID:       GetNextUserID(),
		Username: data.Username,
		IsAdmin:  data.IsAdmin,
	}

	Users = append(Users, user)
	room.Users = append(room.Users, user)

	return user, nil
}
