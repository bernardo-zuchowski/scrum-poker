package gresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	gmodels "multipoker/api/graphql/models"
	"multipoker/internal/dto"
	"multipoker/internal/models"
	"multipoker/internal/repository"
)

func (r *mutationResolver) CreateUser(ctx context.Context, data gmodels.CreateUserInput) (*models.User, error) {
	dto := &dto.CreateUserDTO{
		Username: data.Username,
		RoomID:   data.RoomID,
	}

	return repository.CreateUser(dto)
}

func (r *mutationResolver) CreateVote(ctx context.Context, data gmodels.CreateVoteInput) (*models.User, error) {
	dto := &dto.CreateVoteDTO{
		Vote:   data.Vote,
		UserID: data.UserID,
	}

	return repository.CreateVote(dto)
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return repository.GetAllUsers()
}
