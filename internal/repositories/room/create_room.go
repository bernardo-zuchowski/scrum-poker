package repositories

import (
	"multipoker/internal/dto"
	"multipoker/internal/models"
)

var id int

func getNextID() int {
	id++
	return id
}

func CreateRoom(data *dto.CreateRoomDTO) (*models.Room, error) {
	// user := &models.User{
	// }

	room := &models.Room{
		ID:    getNextID(),
		Title: data.Title,
	}

	return room, nil
}

// func UpdateStatus(id int, *models.User) *models.Room {
// 	for _, todo := range todos {
// 		if todo.ID == id {
// 			todo.Done = status
// 			return todo
// 		}
// 	}

// 	return nil
// }
