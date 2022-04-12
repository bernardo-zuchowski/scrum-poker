package userRepository

import (
	"multipoker/internal/dto"
	"multipoker/internal/models"
)

func CreateUser(data *dto.CreateUserDTO) (*models.User, error) {
	user := &models.User{
		Username: data.Username,
		Vote:     data.Vote,
	}

	return user, nil
}
