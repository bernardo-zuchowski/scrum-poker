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
	var users []*models.User

	user := &models.User{
		ID:       getNextID(),
		Username: *data.HostName,
		IsAdmin:  true,
	}

	users = append(users, user)

	room := &models.Room{
		ID:    getNextID(),
		Title: data.Title,
		Users: users,
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
