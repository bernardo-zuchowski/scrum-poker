package repository

import (
	"multipoker/internal/dto"
	"multipoker/internal/models"
)

func CreateVote(data *dto.CreateVoteDTO) (*models.User, error) {
	user := FindUser(data.UserID)

	user.Vote = data.Vote

	return user, nil
}
