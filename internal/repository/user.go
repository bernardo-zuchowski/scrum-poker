package repository

var userID int

func GetNextUserID() int {
	userID++
	return userID
}
