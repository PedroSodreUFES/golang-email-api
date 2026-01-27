package models

type User struct {
	ID             int32  `json:"id"`
	FullName       string `json:"full_name"`
	PasswordHash   string `json:"password_hash"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}
