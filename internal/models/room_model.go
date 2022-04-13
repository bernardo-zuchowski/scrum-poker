package models

type Room struct {
	ID          int
	Title       string
	HostName    *User
	Users       []*User
	AverageVote float32
}
