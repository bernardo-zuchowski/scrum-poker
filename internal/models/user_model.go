package models

type User struct {
	ID       int
	Username string
	Vote     float32
	IsAdmin  bool
}
