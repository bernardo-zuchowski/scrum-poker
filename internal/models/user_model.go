package models

type User struct {
	ID       int
	Username string
	Vote     *int
	IsVoting bool
	IsAdmin  bool
}
