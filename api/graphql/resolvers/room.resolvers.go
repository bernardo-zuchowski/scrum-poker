package gresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"multipoker/api/graphql/generated"
	gmodels "multipoker/api/graphql/models"
	"multipoker/internal/dto"
	roomRepository "multipoker/internal/repositories/room"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, data gmodels.CreateRoomInput) (*gmodels.Room, error) {
	dto := &dto.CreateRoomDTO{
		Title: data.Title,
	}

	room, err := roomRepository.CreateRoom(dto)

	if err != nil {
		return nil, err
	}

	gromm := &gmodels.Room{
		ID:          room.ID,
		Title:       room.Title,
		AverageVote: room.AverageVote,
	}

	return gromm, nil
}

func (r *mutationResolver) UpdateRoom(ctx context.Context, id int) (*gmodels.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Rooms(ctx context.Context) ([]*gmodels.Room, error) {
	panic(fmt.Errorf("not implemented"))
}
