package repository

import (
	"multipoker/internal/dto"
	"multipoker/internal/models"
)

func CreateRoom(data *dto.CreateRoomDTO) (*models.Room, error) {
	room := &models.Room{
		ID:    GetNextRoomID(),
		Title: data.Title,
	}

	Rooms = append(Rooms, room)

	CreateUser(
		&dto.CreateUserDTO{
			Username: *data.HostName,
			RoomID:   room.ID,
		},
	)

	return room, nil
}
