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

func (r *mutationResolver) CreateRoom(ctx context.Context, data gmodels.CreateRoomInput) (*models.Room, error) {
	dto := &dto.CreateRoomDTO{
		Title:    data.Title,
		HostName: &data.HostName,
	}

	return repository.CreateRoom(dto)
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*models.Room, error) {
	return repository.GetAllRooms()
}

func (r *queryResolver) Votes(ctx context.Context, data gmodels.GetRoomVotesInput) (*models.Room, error) {
	dto := &dto.VotesDTO{
		RoomID: data.RoomID,
	}

	return repository.RevealVotes(dto), nil
}

func (r *roomResolver) AverageVote(ctx context.Context, obj *models.Room) (float64, error) {
	return float64(obj.AverageVote), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type roomResolver struct{ *Resolver }
