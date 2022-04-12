package gresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"multipoker/api/graphql/generated"
	gmodels "multipoker/api/graphql/models"
	"multipoker/internal/dto"
	"multipoker/internal/models"
	userRepository "multipoker/internal/repositories"
)

func (r *mutationResolver) CreateUser(ctx context.Context, data gmodels.CreateUserInput) (*models.User, error) {
	dto := &dto.CreateUserDTO{
		Username: data.Username,
		Vote:     data.Vote,
	}

	return userRepository.CreateUser(dto)
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return userRepository.GetAllUsers()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
