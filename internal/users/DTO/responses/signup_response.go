package responses

type SignUpResponse struct {
	ID             int32  `json:"id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	ProfilePicture *string `json:"profile_picture"`
}
