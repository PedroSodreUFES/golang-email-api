package responses

import "time"

type MyReceivedEmailsResponse struct {
	ID             int32     `json:"id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Wasseen        bool      `json:"wasseen"`
	CreatedAt      time.Time `json:"sent_at"`
	IDReceiver     int32     `json:"id_receiver"`
	IDSender       int32     `json:"id_sender"`
	FullName       string    `json:"sender_full_name"`
	Email          string    `json:"sender_email"`
	ProfilePicture string    `json:"sender_profile_picture"`
}
