package gresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"multipoker/api/graphql/generated"
	gmodels "multipoker/api/graphql/models"
	"multipoker/internal/dto"
	"multipoker/internal/models"
	roomRepository "multipoker/internal/repositories/room"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, data gmodels.CreateRoomInput) (*models.Room, error) {
	dto := &dto.CreateRoomDTO{
		Title:    data.Title,
		HostName: &data.HostName,
	}

	return roomRepository.CreateRoom(dto)
}

func (r *mutationResolver) UpdateRoom(ctx context.Context, id int) (*models.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
