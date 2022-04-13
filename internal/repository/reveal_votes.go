package repository

import (
	"multipoker/internal/dto"
	"multipoker/internal/models"
)

func RevealVotes(data *dto.VotesDTO) *models.Room {
	var votes []float32
	var averageVote float32

	room := FindRoom(data.RoomID)

	for _, user := range room.Users {
		if user.Vote >= 0 {
			votes = append(votes, user.Vote)
		}
	}

	length := len(votes)

	for i := 0; i < length; i++ {
		averageVote = (votes[i] + averageVote)
		if i == length {
			averageVote = averageVote / float32(length)
		}
	}

	room = &models.Room{
		AverageVote: averageVote,
	}

	return room
}
