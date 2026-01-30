package models

import "time"

type Email struct {
	ID         int32     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Wasseen    bool      `json:"wasseen"`
	CreatedAt  time.Time `json:"created_at"`
	IDReceiver int32     `json:"id_receiver"`
	IDSender   int32     `json:"id_sender"`
}
