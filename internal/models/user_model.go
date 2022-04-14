package models

type User struct {
	ID       int
	Username string
	Vote     float64
	IsAdmin  bool
}
