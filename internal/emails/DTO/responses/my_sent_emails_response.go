package responses

import "time"

type MySentEmailsResponse struct {
	ID             int32     `json:"id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Wasseen        bool      `json:"wasseen"`
	CreatedAt      time.Time `json:"sent_at"`
	IDReceiver     int32     `json:"id_receiver"`
	IDSender       int32     `json:"id_sender"`
	FullName       string    `json:"receiver_full_name"`
	Email          string    `json:"receiver_email"`
	ProfilePicture string    `json:"receiver_profile_picture"`
}
