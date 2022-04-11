package gresolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"multipoker/api/graphql/generated"
	gmodels "multipoker/api/graphql/models"
	"multipoker/internal/models"
	"multipoker/internal/repositories"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, data gmodels.CreateTodoInput) (*models.Todo, error) {
	todo := repository.Create(data.Title, data.Done)
	return todo, nil
}

func (r *mutationResolver) UpdateTodoStatus(ctx context.Context, id int, status bool) (*models.Todo, error) {
	todo := repository.UpdateStatus(id, status)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	todos := repository.FindAll()
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var repository = repositories.NewTodoRepository()
