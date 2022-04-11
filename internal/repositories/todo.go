package repositories

import (
	"multipoker/internal/models"
	"time"
)

type TodoRepository struct{}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

var todos = make([]*models.Todo, 0)
var id int

func getNextID() int {
	id++
	return id
}

func (t *TodoRepository) FindAll() []*models.Todo {
	return todos
}

func (t *TodoRepository) Create(title string, done bool) *models.Todo {
	now := time.Now()

	todo := &models.Todo{
		ID:        getNextID(),
		Title:     title,
		Done:      done,
		CreatedAt: now,
		UpdatedAt: now,
	}

	todos = append(todos, todo)

	return todo
}

func (t *TodoRepository) UpdateStatus(id int, status bool) *models.Todo {
	for _, todo := range todos {
		if todo.ID == id {
			todo.Done = status
			return todo
		}
	}

	return nil
}
