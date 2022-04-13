package gresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"multipoker/api/graphql/generated"
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

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return repository.GetAllUsers()
}

func (r *userResolver) Vote(ctx context.Context, obj *models.User) (float64, error) {
	return float64(obj.Vote), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
